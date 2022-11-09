package redis

import (
	"encoding/json"
	"go-redis/models"
	"log"
	"strconv"
	"time"
)

// 查询商铺缓存
func SearchShopById(id string) (string, error) {
	result, err := rdb.Get("cache:shop:" + id).Result()
	if err != nil {
		log.Println("search shop failed err:", err)
	}
	return result, err
}

// 商铺信息写入redis
func SaveShopCache(id string, shop models.Shop) {
	marshal, _ := json.Marshal(shop)
	rdb.Set("cache:shop:"+id, string(marshal), 60*time.Hour)
}

// 商铺信息空值写入redis
func SaveNilCache(id string) {
	rdb.Set("cache:shop:" + id, "", 2 * time.Minute)
}

// 删除商铺缓存
func DeleteShop(id int) {
	shopId := strconv.Itoa(id)
	err := rdb.Del("cache:shop:" + shopId).Err()
	if err != nil {
		log.Println("delete cache failed err:", err)
	}
}

// 获取锁
func TryLock(key string) bool {
	flag, err := rdb.SetNX(key, "1", 10*time.Second).Result()
	if err != nil {
		log.Println("setnx failed err:", err)
	}
	return flag
}

// 释放锁
func ReleaseLock(key string) {
	err := rdb.Del(key).Err()
	if err != nil {
		log.Println("delete lock failed err:", err)
	}
}

// 自增长
func IncreId(keyPrefix, date string) int64 {
	count, err := rdb.Incr("icr:" + keyPrefix + ":" + date).Result()
	if err != nil {
		log.Println("increment id failed err:", err)
	}
	return count
}
