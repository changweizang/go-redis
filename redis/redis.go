package redis

import (
	"github.com/go-redis/redis"
	"log"
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
