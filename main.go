package main

import (
	"go-redis/redis"
	"go-redis/router"
)

func main() {
	redis.InitRedis()
	router.InitRouter()
}
