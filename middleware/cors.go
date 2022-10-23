package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       //允许所有域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, //允许所有请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type","X-Requested-With"}, //必须加载这个"Authorization"
		AllowCredentials: true,
		ExposeHeaders:    []string{"*", "Authorization"},
		MaxAge:           24 * time.Hour,
	})
}
