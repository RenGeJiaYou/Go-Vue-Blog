package model

import (
	"fmt"
	"go-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Category Category `gorm:"foreignkey:Cid"` //将这个模型的主键绑定到外键 Cid 上
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext;not null" json:"content"`
	Image    string   `gorm:"type:varchar(100);not null" json:"image"`
}

// CreateArt 添加文章
func CreateArt(art *Article) int {
	err := db.Create(&art).Error
	if err != nil {
		fmt.Println("添加 user 失败： ", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询某个分类下所有文章

// GetArt todo 获取单篇文章
func GetArt() {

}

// GetArts 获取文章列表
func GetArts(pageSize int, pageNum int) ([]Article, int64) {
	var articles []Article
	var total int64

	err := db.
		Find(&articles).
		Count(&total).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).Error
	if err != nil {
		fmt.Println("查找用户列表失败： ", err)
		return nil, 0
	}
	return articles, total
}

// EditArt 编辑文章
func EditArt(art *Article, id int) int {
	var maps = make(map[string]interface{})
	maps["title"] = art.Title
	maps["cid"] = art.Cid
	maps["desc"] = art.Desc
	maps["content"] = art.Content
	maps["image"] = art.Image

	err := db.
		Model(&art).
		Where("id = ?", id).
		Updates(maps).
		Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.
		Where("id = ?", id).
		Delete(&art).
		Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
