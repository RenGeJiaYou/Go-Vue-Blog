package v1

import (
	"github.com/gin-gonic/gin"
	"go-vue-blog/middleware"
	"go-vue-blog/model"
	"go-vue-blog/utils/errmsg"
	"net/http"
)

// 专做登陆验证
func CheckLogin(c *gin.Context) {
	var data model.User
	var token string
	c.ShouldBindJSON(&data)

	code := model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, _ = middleware.GenerateToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
