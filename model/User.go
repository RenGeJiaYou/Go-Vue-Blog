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
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=2" label:"角色码"` //validate 将 0 视作空值。所以修改为管理员：role=1；游客role=2
}

//model 层的函数签名基本和 Controller 层一致。分工如下：
//		Controller 和前端交互 JSON 数据和请求
//		Model 根据 Controller 所得的 JSON 数据和请求对数据库进行操作

// CheckUser 检查用户是否存在
// Params name string	一个 User 对象的用户名
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

// CheckUpdateUser 更新用户时重名检测，防止更名为其它已存在的用户名
// Params id uint		待更新 User 对象的id
// Params name string	待更新 User 对象的新用户名
// Return code int		一个状态码，表示正确或错误
func CheckUpdateUser(id int, name string) int {
	var u User //作用：1.规定从哪个表查询  2.将查询结果灌注到该对象中
	db.
		Select("id,username").
		Where("username = ?", name).
		First(&u) //字段名全部是小写
	if u.ID == uint(id) || u.ID <= 0 {
		//说明本次更新没有修改用户名 || 更新的用户名未被占用
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_USERNAME_USED
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

// GetUser 获取单个用户
// Params id int	搜索所需的用户ID
// Return u User	所查找的用户
func GetUser(id int) (User, int) {
	var u User

	/*SELECT * FROM user WHERE ID= id*/
	err := db.Limit(1).Where("id = ?", id).Find(&u).Error

	if err != nil {
		return u, errmsg.ERROR_USER_NOT_EXIST
	}
	return u, errmsg.SUCCESS
}

// GetUsers 获取用户列表
// Params username string	搜索所需的用户名
// Params pageSize int		一页的数量
// Params pageNum int		当前的页码
// Return []User			用户列表
// Return int64				用户数量
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var user User
	var users []User
	var total int64

	//非搜索行为
	if username == "" {
		err = db.
			Find(&users).
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).Error
		db.Model(&user).Count(&total)
	}

	//搜索行为（Count()要在 Limit()/Offset() 之前，否则报错）
	err := db.
		Where("username LIKE ?", username+"%").
		Find(&users).
		Count(&total).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).Error

	if err == gorm.ErrRecordNotFound {
		fmt.Println("查找用户列表失败： ", err)
		return nil, 0
	}
	return users, total
}

// EditUser 修改用户
func EditUser(user *User, id int) int {
	//使用map方式,因为 struct 无法更新零值
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role

	err := db.
		Model(&user).
		Where("id = ?", id).
		Updates(maps).
		Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户 不必考虑判空，因为前端一定是在一个条目上删除，不可能空。
// Param id int			用户ID
// Return code int		状态码
func DeleteUser(id int) int {
	var user User
	err := db.
		Where("id = ?", id).
		Delete(&user).
		Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
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

// CheckLogin 登陆验证
func CheckLogin(username, password string) int {
	var user User

	tx := db.
		Where("username = ?", username).
		First(&user)

	if tx.RowsAffected == 0 { /* if user.ID == 0 */
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	// 1:管理员 2:游客
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}

	return errmsg.SUCCESS
}
