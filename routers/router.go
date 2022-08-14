package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-vue-blog/api/v1"
	"go-vue-blog/utils"
)

// InitRouter router 入口文件
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//user 的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		router.GET()
		//article 的路由接口

		//category 的路由接口
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
