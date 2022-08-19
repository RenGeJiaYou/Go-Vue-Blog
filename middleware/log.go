package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/govueblog.log"

	//打开指定处的文件，并指定权限为：可读可写，可创建
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755) //0755-> rwx r-x r-x linux知识
	if err != nil {
		fmt.Println("err:", err)
	}

	log := logrus.New()
	log.Out = src
	return func(c *gin.Context) {
		// 一.配置所需的 Fields
		startTime := time.Now()
		c.Next()
		spendTime := time.Since(startTime).Milliseconds()
		ST := fmt.Sprintf("%d ms", spendTime) // 1.API 调用耗时
		hostName, err := os.Hostname()        // 2.主机名
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()    // 3.状态码
		clientIP := c.ClientIP()           // 4.请求客户端的 IP
		userAgent := c.Request.UserAgent() // 5.用户代理，通常是某个浏览器。dev环境下是apipost
		dataSize := c.Writer.Size()        // 6.响应报文 body 的字节长度
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method   // 7.请求方法
		path := c.Request.RequestURI // 8.请求 URL

		// 二.从标准记录器创建一个条目，并向其中添加多个字段(隐式添加 log 本身的时间戳,信息等 fields )
		entry := log.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": ST,
			"IP":        clientIP,
			"UserAgent": userAgent,
			"Method":    method,
			"DataSize":  dataSize,
			"Path":      path,
		})

		// Errors 保存了使用当前context的所有中间件/handler 所产生的全部错误信息。
		// 源码注释： Errors is a list of errors attached to all the handlers/middlewares who used this context.
		// 三.将系统内部的错误 log 出去
		if len(c.Errors) > 0 {
			log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		// 四.根据状态码决定打印 log 的等级
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
