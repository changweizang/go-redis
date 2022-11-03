package redis

import (
	"encoding/json"
	"go-redis/models"
	"log"
	"time"
)

// 查询商铺缓存
func SearchShopById(id string) string {
	result, err := rdb.Get("cache:shop:" + id).Result()
	if err != nil {
		log.Println("search shop failed err:", err)
	}
	return result
}

// 商铺信息写入redis
func SaveShopCache(id string, shop models.Shop) {
	marshal, _ := json.Marshal(shop)
	rdb.Set("cache:shop:"+id, string(marshal), 60*time.Hour)
}

// 删除商铺缓存
func DeleteShop(id string) {
	err := rdb.Del("cache:shop:" + id).Err()
	if err != nil {
		log.Println("delete cache failed err:", err)
	}
}

