package api

import "github.com/gin-gonic/gin"

func QueryByIdHandle(ctx *gin.Context) {
	// 从redis查询商铺缓存
	// 存在，直接返回
	// 不存在，根据id查询数据库
	// 不存在，返回错误
	// 存在，写入redis
	// 返回结果
}