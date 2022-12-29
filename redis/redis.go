package redis

import (
	redission "github.com/changweizang/go-redisson"
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client
var c *redission.Common

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "106.15.199.75:6379",
		Password: "123456",
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	c = redission.InitRlock(rdb)
}

func RedisClient() *redis.Client {
	return rdb
}

func InitRlock() *redission.Common {
	return c
}
