package api

import (
	"go-redis/models"
	"go-redis/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// post
// input voucher
// 添加秒杀优惠券
func AddSkillVoucher(ctx *gin.Context) {
	res := utils.InitResBody()
	req := utils.ReqSkillVoucher{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Code = http.StatusBadRequest
		res.Message = "解析参数失败"
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	// 添加到秒杀券和优惠券两张表中
	voucher := models.Voucher{}
	voucher.ShopId = req.ShopId
	voucher.Title = req.Title
	voucher.SubTitle = req.SubTitle
	voucher.Rules = req.Rules
	voucher.PayValue = req.PayValue
	voucher.ActualValue = req.ActualValue
	voucher.Type = req.Type
	voucher.CreateTime = time.Now()
	voucher.UpdateTime = time.Now()
	err, VoucherId := models.AddVoucher(voucher)
	if err != nil {
		log.Println("save voucher failed err:", err)
		res.Code = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, res)
		return	
	}
	skillVoucher := models.SeckillVoucher{}
	skillVoucher.VoucherId = VoucherId
	skillVoucher.Stock = req.Stock
	skillVoucher.BeginTime = req.BeginTime
	skillVoucher.EndTime = req.EndTime
	skillVoucher.CreateTime = time.Now()
	skillVoucher.UpdateTime = time.Now()
	if err := models.AddSkillVoucher(skillVoucher); err != nil {
		log.Println("save voucher failed err:", err)
		res.Code = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.Message = "添加秒杀优惠券成功"
	ctx.JSON(http.StatusOK, res)
}
