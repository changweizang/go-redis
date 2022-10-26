package router

import (
	"github.com/gin-gonic/gin"
	"go-redis/api"
)

func initRouter() {
	router := gin.Default()
	router.POST("/api/code", api.CodeHandle)
}
