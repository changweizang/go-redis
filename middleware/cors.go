package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	})
}
