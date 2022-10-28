package main

import (
	"go-redis/models"
	"go-redis/redis"
	"go-redis/router"
)

func main() {
	models.InitMysql()
	redis.InitRedis()
	router.InitRouter()
}
