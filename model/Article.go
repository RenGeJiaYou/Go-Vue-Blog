package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Category Category `gorm:"foreignkey:Cid"`	//将这个模型的主键绑定到外键 Cid 上
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext;not null" json:"content"`
	Image    string   `gorm:"type:varchar(100);not null" json:"image"`
}
