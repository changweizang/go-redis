package api

import (
	"github.com/gin-gonic/gin"
	"go-redis/models"
	"go-redis/redis"
	"go-redis/utils"
	"net/http"
)

func QueryByIdHandle(ctx *gin.Context) {
	res := utils.ResBody{}
	id := ctx.Param("id")
	// 从redis查询商铺缓存
	cacheShop := redis.SearchShopById(id)
	// 存在，直接返回
	if cacheShop != "" {
		res.Code = http.StatusOK
		res.Message = "查询到信息"
		res.Data = cacheShop
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 不存在，根据id查询数据库
	shop := models.SearchShopById(id)
	// 写入redis
	redis.SaveShopCache(id, shop)
	// 返回结果
	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = shop
	ctx.JSON(http.StatusOK, res)
}