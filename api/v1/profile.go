package v1

import (
	"github.com/gin-gonic/gin"
	"go-vue-blog/model"
	"go-vue-blog/utils/errmsg"
	"net/http"
	"strconv"
)

// GetProfile 获取个人简介
func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pid := uint8(id)

	profile, code := model.GetProfile(pid)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    profile,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateProfile 更新个人简介
func UpdateProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pid := uint8(id)

	var data model.Profile
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateProfile(pid, &data) //传入的是一个较大的对象时，应该用传址调用

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
