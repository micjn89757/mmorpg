package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			// AllowOrigins
			AllowOrigins: []string{"*"}, // 等同于允许所有域名
			AllowMethods: []string{"GET", "POST"},
			AllowHeaders: []string{"*"}, // 允许的请求头
			AllowCredentials: true,
			ExposeHeaders: []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			MaxAge: 12 * time.Hour,  // 预检请求结果可以缓存多久
		},
	)
}