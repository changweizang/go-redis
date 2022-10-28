package api

import (
	"github.com/gin-gonic/gin"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"
)

// LoginHandle /*
/*
post
input {phone code}
校验验证码
判断用户是否存在
不存在：创建新用户
保存用户信息到redis
 */
func LoginHandle(ctx *gin.Context) {
	res := utils.ResBody{}
	loginBody := utils.Login{}
	err := ctx.ShouldBindJSON(&loginBody)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = "解析参数失败"
		log.Fatalln(err)
	}
	// 校验验证码
	ok, err := redis.CheckLoginCode(loginBody.Phone, loginBody.Code)
	if err != nil {
		res.Code = http.StatusOK
		res.Message = err.Error()
		return
	}
	if !ok {
		res.Code = http.StatusOK
		res.Message = "验证码错误"
		return
	}

}
