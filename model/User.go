package model

import (
	"fmt"
	"go-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

//model 层的函数签名基本和 Controller 层一致。分工如下：
//		Controller 和前端交互 JSON 数据和请求
//		Model 根据 Controller 所得的 JSON 数据和请求对数据库进行操作

// CheckUser 检查用户是否存在
// Params user *User	一个填充了前端 JSON 数据的 User 对象
// Return code int		一个状态码，表示正确或错误
func CheckUser(name string) int {
	var u User //作用：1.规定从哪个表查询  2.将查询结果灌注到该对象中
	db.
		Select("id").
		Where("username = ?", name).
		First(&u) //字段名全部是小写
	if u.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS

}

// CreateUser 向数据库添加用户
// Params user *User	一个填充了前端 JSON 数据的 User 对象
// Return code int		一个状态码，表示正确或错误
func CreateUser(user *User) int { // Go 的 struct 是引用类型，作为参数传入时必须作为一个指针
	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("添加 user 失败： ", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err := db.
		Find(&users).
		Count(&total).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).Error
	if err != nil {
		fmt.Println("查找用户列表失败： ", err)
		return nil, 0
	}
	return users, total

}
