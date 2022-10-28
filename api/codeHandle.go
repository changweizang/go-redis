package api

import (
	"github.com/gin-gonic/gin"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"
)

// post
// input: phone
// 生成验证码并存入redis
func CodeHandle(ctx *gin.Context) {
	phoneBody := utils.PhoneBody{}
	res := utils.ResBody{}
	err := ctx.ShouldBind(&phoneBody)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = "获取号码失败"
		log.Println(err)
		ctx.JSON(http.StatusOK, res)
	}
	// 获取验证码
	code := utils.GetRandomCode()
	log.Println("验证码：", code)
	// 保存到redis
	err = redis.SavePhoneCode(phoneBody.Phone, code)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = "redis错误"
		log.Println(err)
		ctx.JSON(http.StatusOK, res)
	}
	res.Code = http.StatusOK
	res.Message = "获取验证码成功"
	res.Data = code
	ctx.JSON(http.StatusOK, res)
}
