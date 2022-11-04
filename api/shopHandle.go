package api

import (
	"go-redis/models"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get
// input id
func QueryShopByIdHandle(ctx *gin.Context) {
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

// put
// input shop
func UpdateShopHandle(ctx *gin.Context) {
	res := utils.ResBody{}
	shop := models.Shop{}
	err := ctx.ShouldBindJSON(&shop)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = "解析参数失败"
		log.Println(err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	shopId := shop.Id
	if shopId <= 0 {
		res.Code = http.StatusBadRequest
		res.Message = "店铺id错误"
		ctx.JSON(http.StatusOK, res)
	}
	// 更新数据库中商铺信息
	models.UpdateShop(shop)
	// 删除该条缓存，下次查询时再添加缓存
	redis.DeleteShop(shop.Id)
	res.Code = http.StatusOK
	res.Message = "删除缓存成功"
	ctx.JSON(http.StatusOK, res)
}
