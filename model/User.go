package model

import (
	"encoding/base64"
	"fmt"
	"go-vue-blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
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

// GetUsers 获取用户列表
// Params pageSize int	一页的数量
// Params pageNum int	当前的页码
// Return []User		用户列表
// Return int64			用户数量
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

// ScryptPw 在存入数据库前对用户密码加密，将被 一个钩子函数调用
// Param password string	用户原始密码
// Return hash string
func ScryptPw(password string) string {
	const keyLen = 16
	salt := []byte{53, 234, 5, 64, 23, 97, 34, 4, 77, 31, 56, 92}

	HashPw, err := scrypt.Key([]byte(password), salt, 1<<12, 18, 1, keyLen)
	if err != nil {
		fmt.Println("对明文密码加密出错：", err)
	}
	//base64 提供 bit -> 字节 的转换
	return base64.StdEncoding.EncodeToString(HashPw)
}

// BeforeSave 钩子函数，在数据存入数据库前哈希密码
func (u *User) BeforeSave(_ *gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	return nil
}
