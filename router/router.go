package router

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-redis/api"
	"go-redis/middleware"
)

func ListenPort() string {
	var port string
	flag.StringVar(&port, "p", "9090", "端口号 默认9090")
	flag.Parse()
	return fmt.Sprintf("%v", port)
}

func InitRouter(port string) {
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

	router.Run(":" + port)
}
