package api

import (
	"github.com/gin-gonic/gin"
	"go-redis/models"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

// post
// input voucherId
// 秒杀优惠券
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
	// 5.扣减库存
	err = models.DecVoucherSock(seckillVoucher)
	if err != nil {
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
	token := ctx.Request.Header.Get("Authorization")
	userId := redis.GetUser(token)
	atoiUserId, _ := strconv.Atoi(userId)
	voucherOrder.UserId = atoiUserId
	// 6.3.代金券id
	atoiVoucherId, _ := strconv.Atoi(voucherId)
	voucherOrder.VoucherId = atoiVoucherId
	// 6.4.保存订单信息
	voucherOrder.CreateTime = time.Now()
	voucherOrder.UpdateTime = time.Now()
	err = models.SaveVoucherOrder(voucherOrder)
	if err != nil {
		log.Println("save voucher order failed err:", err)
		res.Message = "保存订单失败"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 7.返回订单id
	res.Data = orderId
	ctx.JSON(http.StatusOK, res)
	return
}
