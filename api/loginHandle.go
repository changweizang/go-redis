package api

import (
	"go-redis/models"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		log.Println(err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 校验验证码
	ok, err := redis.CheckLoginCode(loginBody.Phone, loginBody.Code)
	if err != nil {
		res.Code = http.StatusOK
		res.Message = err.Error()
		log.Println(err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	if !ok {
		res.Code = http.StatusOK
		res.Message = "验证码错误"
		log.Println(err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 用户状态
	if exist := models.IfUserExist(loginBody.Phone); !exist {
		err := models.AddUser(loginBody.Phone)
		if err != nil {
			res.Code = http.StatusOK
			res.Message = "添加用户失败"
			log.Println(err)
			ctx.JSON(http.StatusOK, res)
			return
		}
	}
	// 查询用户
	user := models.SearchUserByPhone(loginBody.Phone)
	// 随机生成token作为登录令牌
	token := utils.GetUUid()
	// 将user对象转化为hash存储并设置有效期
	redis.SaveUser(token, user)
	// 返回token
	res.Code = http.StatusOK
	res.Message = "登录成功"
	res.Data = token
	ctx.JSON(http.StatusOK, res)
}
