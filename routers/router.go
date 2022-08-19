package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-vue-blog/api/v1"
	"go-vue-blog/middleware"
	"go-vue-blog/utils"
)

// InitRouter router 入口文件
func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//user 的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		//article 的路由接口
		auth.POST("article/add", v1.AddArt)      //添加文章
		auth.PUT("article/:id", v1.EditArt)      //编辑文章
		auth.DELETE("article/:id", v1.DeleteArt) //删除文章

		//category 的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCate)

		//文件上传的路由接口
		auth.POST("upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{
		//user 的路由接口
		router.GET("users", v1.GetUsers)
		router.POST("user/add", v1.AddUser)

		//article 的路由接口
		router.GET("articles", v1.GetArts)             //获取全部文章
		router.GET("article/info/:id", v1.GetArt)      //获取单篇文章
		router.GET("article/list/:cid", v1.GetCateArt) //获取指定分类下的文章

		//category 的路由接口
		router.GET("categories", v1.GetCategories)

		//登录的接口
		router.POST("login", v1.CheckLogin)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
