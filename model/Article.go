package model

import (
	"fmt"
	"go-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

// ORM 中，总要将 Where() 放在 Find()/First() 前。否则将因为过早查询而略过了条件

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

// GetCateArt todo 查询某个分类下所有文章
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArt []Article
	var total int64
	err := db.
		Preload("Category").
		Limit(pageSize).
		Offset((pageNum-1)*pageSize).
		Where("cid = ?", cid).
		Find(&cateArt).
		Error

	//记录总数。Model() 不仅仅能传已经定义好的go struct 对象，也包括某次查询的视图
	db.
		Model(&cateArt).
		Where("cid =?", cid).
		Count(&total)

	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArt, errmsg.SUCCESS, total

}

// GetArt  获取单篇文章
func GetArt(id int) (Article, int) {
	var art Article

	//Where() 应该在 Find() 前。否则将先进行查询，而 Where() 条件未能起作用
	err := db.
		Preload("Category"). //若不预加载 Category ，返回的Article 的 Category 数据子项将是空值
		Where("id = ?", id).
		Find(&art).
		Error
	if err != nil {
		fmt.Println("获取单篇文章 在 ORM 阶段失败")
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS

}

// GetArts 获取文章列表
func GetArts(pageSize int, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var total int64

	err := db.
		Preload("Category").
		Find(&articles).
		Count(&total).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).Error
	if err != nil {
		fmt.Println("查找用户列表失败： ", err)
		return nil, errmsg.ERROR_ART_NOT_EXIST, 0
	}
	return articles, errmsg.SUCCESS, total
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
