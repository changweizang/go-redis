package api

import (
	"encoding/json"
	"go-redis/models"
	"go-redis/redis"
	"go-redis/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// get
// input id
// 缓存穿透/缓存击穿处理
func QueryShopByIdHandle(ctx *gin.Context) {
	res := utils.ResBody{}
	id := ctx.Param("id")

	// 缓存穿透
	if 1 == 2 {
		queryShopPassThrough(ctx)
	}
	// 缓存击穿
	queryShopMutex(ctx, res, id)
}

// 缓存击穿
func queryShopMutex(ctx *gin.Context, res utils.ResBody, id string) error {
	// 从redis查询商铺缓存
	cacheShop, err := redis.SearchShopById(id)
	// 存在，直接返回
	if cacheShop != "" {
		shoper := models.Shop{}
		if err2 := json.Unmarshal([]byte(cacheShop), &shoper); err2 != nil {
			log.Println("unmarshal json failed err:", err2)
			return nil
		}
		res.Code = http.StatusOK
		res.Message = "查询到信息"
		res.Data = shoper
		ctx.JSON(http.StatusOK, res)
		return nil
	}
	// 查询到的缓存为空，直接返回空值，避免恶意访问
	if err == nil && cacheShop == "" {
		res.Code = http.StatusOK
		res.Message = "店铺信息不存在"
		ctx.JSON(http.StatusOK, res)
		return nil
	}
	// 获取互斥锁
	lockKey := "lock:shop:" + id
	flag := redis.TryLock(lockKey)
	// 判断是否获取成功
	// 失败则休眠重试
	if !flag {
		time.Sleep(50 * time.Millisecond)
		return queryShopMutex(ctx, res, id)
	}

	// 成功则根据id查询数据库
	shop, err := models.SearchShopById(id)
	// 模拟重建的延时
	//time.Sleep(200 * time.Millisecond)
	// 不存在，将空值写入redis，返回错误, 避免用户多次恶意访问
	if err != nil {
		redis.SaveNilCache(id)
		res.Code = http.StatusOK
		res.Message = "店铺不存在"
		ctx.JSON(http.StatusOK, res)
		return nil
	}
	// 存在写入redis
	redis.SaveShopCache(id, shop)
	// 释放互斥锁
	redis.ReleaseLock(lockKey)
	// 返回结果
	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = shop
	ctx.JSON(http.StatusOK, res)
	return nil
}

// 缓存穿透
func queryShopPassThrough(ctx *gin.Context) {
	res := utils.ResBody{}
	id := ctx.Param("id")
	// 从redis查询商铺缓存
	cacheShop, err := redis.SearchShopById(id)
	shoper := models.Shop{}
	if err2 := json.Unmarshal([]byte(cacheShop), &shoper); err2 != nil {
		log.Println("unmarshal json failed err:", err2)
		return
	}
	// 存在，直接返回
	if cacheShop != "" {
		res.Code = http.StatusOK
		res.Message = "查询到信息"
		res.Data = shoper
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 查询到的缓存为空，直接返回空值，避免恶意访问
	if err == nil && cacheShop == "" {
		res.Code = http.StatusOK
		res.Message = "店铺信息不存在"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 未查询到缓存，根据id查询数据库
	shop, err := models.SearchShopById(id)
	// 不存在，将空值写入redis，返回错误, 避免用户多次恶意访问
	if err != nil {
		redis.SaveNilCache(id)
		res.Code = http.StatusOK
		res.Message = "店铺不存在"
		ctx.JSON(http.StatusOK, res)
		return
	}
	// 存在写入redis
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
