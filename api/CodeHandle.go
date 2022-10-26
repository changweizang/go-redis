package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// post
// 生成验证码并存入redis
func CodeHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello ")
}
