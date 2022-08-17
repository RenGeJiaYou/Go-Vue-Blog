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

		//article 的路由接口
		router.POST("article/add", v1.AddArt)          //添加文章
		router.GET("articles", v1.GetArts)             //获取全部文章
		router.GET("article/info/:id", v1.GetArt)      //获取单篇文章
		router.GET("article/list/:cid", v1.GetCateArt) //获取指定分类下的文章
		router.PUT("article/:id", v1.EditArt)          //编辑文章
		router.DELETE("article/:id", v1.DeleteArt)     //删除文章

		//category 的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("categories", v1.GetCategories)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCate)

	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
