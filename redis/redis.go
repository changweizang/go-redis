package redis

import (
	"encoding/json"
	"go-redis/models"
	"log"
	"time"

	"github.com/fatih/structs"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "106.15.199.75:6379",
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
}

func RedisClient() *redis.Client {
	return rdb
}

// 保存验证码
func SavePhoneCode(phone, code string) error {
	return rdb.Set("phone:" + phone, code, 10*time.Minute).Err()
}

// 校验验证码
func CheckLoginCode(phone, code string) (bool, error) {
	result, err := rdb.Get("phone:" + phone).Result()
	if err != nil {
		return false, err
	}
	if result == code {
		return true, nil
	}
	return false, nil
}

// 将user存入redis
func SaveUser(token string, user models.User) {
	m := structs.Map(&user)
	rdb.HMSet("login:" + token, m)
	rdb.Expire("login:" + token, 60 * time.Hour)
}

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
	rdb.Set("cashe:shop:" + id, string(marshal), 60 * time.Hour)
}
