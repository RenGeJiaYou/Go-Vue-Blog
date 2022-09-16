package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-vue-blog/model"
	"go-vue-blog/utils/errmsg"
	"go-vue-blog/utils/validator"
	"net/http"
	"strconv"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("AddUser() 绑定 JSON 出错：", err)
	}

	msg, errCode := validator.Validator(&user)
	if errCode != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"data":    msg,
			"message": errmsg.GetErrMsg(errCode),
		})
		return
	}
	//检查重名
	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
	}

	//返回状态:回复http.StatusOK是为了说明 web 服务正常执行了
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})

}

//查询单个用户（在 Blog 系统中用处不大）

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	//从请求报文的 params 提取数据
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	if pageSize == 0 {
		pageSize = -1 //GORM 的Limit(-1) 表示不要 Limit() 这个限制
	}
	if pageNum == 0 {
		pageNum = 1 //GORM 的Offset((1-1)*pageSize) 表示不要 Offset() 这个限制
	}

	//查询数据库
	users, total := model.GetUsers(username, pageSize, pageNum)

	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    users,
		"total":   total,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user model.User
	c.ShouldBindJSON(&user)
	code := model.CheckUser(user.Username)

	if code == errmsg.SUCCESS {
		model.EditUser(&user, id) //实际上只能改 username 和 role
	}
	if code == errmsg.ERROR_USERNAME_USED {
		// 若重名，不再调用后续的函数处理
		fmt.Println("ERROR_USERNAME_USED")
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"delete": "i am delete",
	})
}
