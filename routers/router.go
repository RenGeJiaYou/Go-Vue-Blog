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
	//先设置setMode(),再创建 gin.New() 才能正确运行配置
	gin.SetMode(utils.AppMode)

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//前端资源托管(pattern 参数 从项目根目录开始)
	r.LoadHTMLGlob("static/admin/index.html")
	//第1个参数是路由的路径，第2个参数是文件的路径
	r.Static("admin/static", "static/admin/static")
	//在浏览器地址栏输入localhost:<socket>/admin 即可访问Vue Router设置 path 为 '/' 的 component
	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)

	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//user 的路由接口
		auth.PUT("user/:id", v1.EditUser) //:id 意为 前端axios会把“user/”后面的任何一个数据视为c.Param 的一个键值对，并且 key 为id，value为实际的数据（实际传入该数据时不用加":")
		auth.PUT("user/resetpw/:id", v1.ResetPass)
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

		//修改个人简介的接口
		auth.PUT("profile/:id", v1.UpdateProfile)
	}

	router := r.Group("api/v1")
	{
		//user 的路由接口
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo)
		router.POST("user/add", v1.AddUser)

		//article 的路由接口
		router.GET("articles", v1.GetArts)             //获取全部文章
		router.GET("article/info/:id", v1.GetArt)      //获取单篇文章
		router.GET("article/list/:cid", v1.GetCateArt) //获取指定分类下的文章

		//category 的路由接口
		router.GET("categories", v1.GetCategories)

		//登录的接口
		router.POST("login", v1.CheckLogin)

		//获取个人简介的接口
		router.GET("profile/:id", v1.GetProfile)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
