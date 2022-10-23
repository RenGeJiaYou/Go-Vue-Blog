package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-vue-blog/model"
	"go-vue-blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var cate model.Category
	if err := c.ShouldBindJSON(&cate); err != nil {
		fmt.Println("AddCategory() 绑定 JSON 出错：", err)
	}
	//检查重名
	code := model.CheckCate(cate.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&cate)
	}

	//返回状态:回复http.StatusOK是为了说明 web 服务正常执行了
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cate,
		"message": errmsg.GetErrMsg(code),
	})

}

//查询单个分类下的文章

// GetCategories 查询分类列表
func GetCategories(c *gin.Context) {
	//从请求报文的 params 提取数据
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	name := c.Query("name")

	if pageSize == 0 {
		pageSize = -1 //GORM 的Limit(-1) 表示不要 Limit() 这个限制
	}
	if pageNum == 0 {
		pageNum = 1 //GORM 的Offset((1-1)*pageSize) 表示不要 Offset() 这个限制
	}

	//查询数据库
	cates, total := model.GetCategories(name, pageSize, pageNum)

	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    cates,
		"total":   total,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})
}

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var cate model.Category
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		fmt.Println("c.ShouldBindJSON(&cate) 失败")
	}
	code := model.CheckCate(cate.Name)

	if code == errmsg.SUCCESS {
		model.EditCate(&cate, id)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		// 若重名，不再调用后续的函数处理
		fmt.Println("ERROR_CATENAME_USED")
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCate 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"delete":  "i am delete",
	})
}
