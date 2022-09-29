package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-vue-blog/model"
	"go-vue-blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddArt 添加文章
func AddArt(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		fmt.Println("AddArt() 绑定 JSON 出错：", err)
	}

	code := model.CreateArt(&article)

	//返回状态:回复http.StatusOK是为了说明 web 服务正常执行了
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetCateArt  查询某个分类下所有文章
func GetCateArt(c *gin.Context) {
	//从请求报文的 params 提取数据
	cid, _ := strconv.Atoi(c.Param("cid")) // 巩固Param() 和 Query() 的区别
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	fmt.Println(cid, pageSize, pageNum)
	if pageSize == 0 {
		pageSize = -1 //GORM 的Limit(-1) 表示不要 Limit() 这个限制
	}
	if pageNum == 0 {
		pageNum = 1 //GORM 的Offset((1-1)*pageSize) 表示不要 Offset() 这个限制
	}
	// 调用 Model 层
	arts, code, total := model.GetCateArt(cid, pageSize, pageNum)
	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    arts,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArt 查询单个文章
func GetArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.GetArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArts 查询文章列表
func GetArts(c *gin.Context) {
	//从请求报文的 params 提取数据
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	if pageSize == 0 {
		pageSize = -1 //GORM 的Limit(-1) 表示不要 Limit() 这个限制
	}
	if pageNum == 0 {
		pageNum = 1 //GORM 的Offset((1-1)*pageSize) 表示不要 Offset() 这个限制
	}

	//查询数据库
	arts, code, total := model.GetArts(title, pageSize, pageNum)

	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    arts,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditArt 编辑文章
func EditArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var art model.Article
	err := c.ShouldBindJSON(&art)
	if err != nil {
		fmt.Println(err)
	}

	code := model.EditArt(&art, id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteArt 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"delete":  "i am delete",
	})
}
