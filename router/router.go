package router

import (
	"github.com/gin-gonic/gin"
	"go-redis/api"
)

func InitRouter() {
	router := gin.Default()
	router.POST("/user/code", api.CodeHandle)
	router.POST("/user/login", api.LoginHandle)
	router.GET("/shop/:id", api.QueryShopByIdHandle)
	router.PUT("/shop", api.UpdateShopHandle)

	router.Run(":8081")
}
