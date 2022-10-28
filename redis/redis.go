package redis

import (
	"github.com/go-redis/redis"
	"log"
	"time"
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
	err := rdb.Set("phone:" + phone, code, 10*time.Minute).Err()
	return err
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
