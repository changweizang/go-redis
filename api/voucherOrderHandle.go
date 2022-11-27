package api

import (
	"fmt"
	"go-redis/models"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var mutex sync.Mutex

// post
// input voucherId
// 秒杀优惠券
// 乐观锁解决库存超卖问题
func SeckillVoucher(ctx *gin.Context) {
	res := utils.InitResBody()
	voucherId := ctx.Param("id")
	// 1.查询优惠券
	seckillVoucher, err := models.QuerySeckillVoucherById(voucherId)
	if err != nil {
		log.Println("query seckillVoucher failed err:", err)
		res.Message = "failed"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 2.判断秒杀是否开始
	seckillBeginTime, err := time.Parse(utils.TimeFormat, seckillVoucher.BeginTime.String())
	if err != nil || time.Now().Before(seckillBeginTime) {
		log.Println("time is too early")
		res.Message = "未到优惠券使用时间"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 3.判断秒杀是否结束
	seckillEndTime, err := time.Parse(utils.TimeFormat, seckillVoucher.EndTime.String())
	if err != nil || seckillEndTime.Before(time.Now()) {
		log.Println("time is too late")
		res.Message = "优惠券过期"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 4.判断库存是否充足
	if seckillVoucher.Stock < 1 {
		log.Println("秒杀优惠券库存不足")
		res.Message = "秒杀优惠券库存不足"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 一人一单
	mutex.Lock()
	defer mutex.Unlock()
	phoneValue, ok := ctx.Get("phone")
	if !ok {
		res.Message = "未获取到用户信息"
		log.Println("未获取到用户信息")
		ctx.JSON(http.StatusOK, res)
		return
	}
	phone := fmt.Sprintf("%v", phoneValue)
	user := models.SearchUserByPhone(phone)
	userId := user.Id
	count, err := models.QueryCountByUserId(userId)
	if err != nil {
		res.Message = "查询订单错误"
		log.Println("查询订单错误")
		ctx.JSON(http.StatusOK, res)
		return
	}
	if count > 0 {
		res.Message = "用户已经购买过一次了"
		log.Println("用户已经购买过一次了")
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 5.开启事务，扣减库存
	tx := models.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 乐观锁处理，扣减库存时判断库存是否大于0
	countStock, err := models.DecVoucherSock(seckillVoucher, tx)
	if err != nil || countStock == 0 {
		tx.Rollback()
		log.Println("update count failed")
		res.Message = "扣减库存失败"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 6.创建订单
	voucherOrder := models.VoucherOrder{}
	// 6.1.订单id
	orderId := redis.NextId("voucherOrder")
	voucherOrder.Id = int(orderId)
	// 6.2.用户id
	voucherOrder.UserId = userId
	// 6.3.代金券id
	atoiVoucherId, _ := strconv.Atoi(voucherId)
	voucherOrder.VoucherId = atoiVoucherId
	// 6.4.保存订单信息
	voucherOrder.CreateTime = time.Now()
	voucherOrder.UpdateTime = time.Now()
	err = models.SaveVoucherOrder(voucherOrder, tx)
	if err != nil {
		tx.Rollback()
		log.Println("save voucher order failed err:", err)
		res.Message = "保存订单失败"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 结束事务
	tx.Commit()
	// 7.返回订单id
	res.Data = orderId
	ctx.JSON(http.StatusOK, res)
	return
}
