package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:  []string{"*"}, //AllowAllOrigins: true,//允许跨域所有网站
			AllowMethods:  []string{"*"}, // 允许的请求方法
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			//AllowCredentials: true, // 是否发送cookie请求
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//},
			MaxAge: 12 * time.Hour,
		})

	}
}
