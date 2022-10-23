package model

import (
	"fmt"
	"go-vue-blog/utils/errmsg"
)

type Category struct {
	ID   uint8  `gorm:"primaryKey" json:"id"`
	Name string `·gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCate 检查标签是否存在
// Params name string	一个 Category 对象的标签名
// Return code int		一个状态码，表示正确或错误
func CheckCate(name string) int {
	var cate Category //作用：1.规定从哪个表查询  2.将查询结果灌注到该对象中
	db.
		Select("id").
		Where("name = ?", name).
		First(&cate) //字段名全部是小写
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS

}

// CreateCate 向数据库添加标签
// Params cate *Cate	一个填充了前端 JSON 数据的 Cate 对象
// Return code int		一个状态码，表示正确或错误
func CreateCate(cate *Category) int { // Go 的 struct 是引用类型，作为参数传入时必须作为一个指针
	err := db.Create(&cate).
		Error
	if err != nil {
		fmt.Println("添加 标签 失败： ", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//todo 查询分类下的所有文章

// GetCategories 获取标签列表
// Params pageSize int	一页的数量
// Params pageNum int	当前的页码
// Return []User		标签列表
// Return int64			标签数量
func GetCategories(name string, pageSize int, pageNum int) ([]Category, int64) {
	var cas []Category
	var total int64
	var err error

	if name == "" {
		err = db.
			Find(&cas).
			Count(&total).
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).Error
	} else {
		err = db.
			Where("name LIKE ?", name+"%").
			Find(&cas).
			Count(&total).
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).Error

	}

	if err != nil {
		fmt.Println("查找标签列表失败： ", err)
		return nil, 0
	}
	return cas, total
}

// EditCate 修改标签
func EditCate(cate *Category, id int) int {
	//使用map方式,因为 struct 无法更新零值
	var maps = make(map[string]interface{})
	maps["name"] = cate.Name

	err := db.
		Model(&cate).
		Where("id = ?", id). //Where() 应该在 Update()前面，否则报错“无Where()条件 ”
		Updates(maps).
		Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate 删除标签 不必考虑判空，因为前端一定是在一个条目上删除，不可能空。
// Param id int			标签ID
// Return code int		状态码
func DeleteCate(id int) int {
	var cate Category
	err := db.
		Where("id = ?", id).
		Delete(&cate).
		Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
