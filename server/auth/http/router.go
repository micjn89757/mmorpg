package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

func initRouter(middleware ...gin.HandlerFunc) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	// 设置新人网络 []string
	r := gin.New()
	r.SetTrustedProxies(nil)

	// 添加中间件
	r.Use(gin.Recovery())
	for _, m := range middleware {
		r.Use(m)
	}

	r.POST("login", Login)
	r.POST("register", Register)
	return r
}