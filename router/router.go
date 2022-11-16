package router

import (
	"github.com/gin-gonic/gin"
	"go-redis/api"
	"go-redis/middleware"
)

func InitRouter() {
	router := gin.Default()
	r := router.Group("/user")
	{
		r.POST("/code", api.CodeHandle)
		r.POST("/login", api.LoginHandle)
	}

	auth := router.Group("/auth")
	auth.Use(middleware.JwtToken())
	{
		auth.GET("/shop/:id", api.QueryShopByIdHandle)
		auth.PUT("/shop", api.UpdateShopHandle)
		auth.POST("/voucher/seckill", api.AddSkillVoucher)
		auth.POST("/voucher-order/seckill/:id", api.SeckillVoucher)
	}

	router.Run(":9090")
}
