package model

import "go-vue-blog/utils/errmsg"

type Profile struct {
	ID     uint8  `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name"`    //用户名
	Img    string `gorm:"type:varchar(100)" json:"img"`    //背景图
	Avatar string `gorm:"type:varchar(100)" json:"avatar"` //头像
	State  string `gorm:"type:varchar(100)" json:"state"`  //职业状态
	Desc   string `gorm:"type:varchar(200)" json:"desc"`   //个人简介
	//若干联系方式
	Qqchat string `gorm:"type:varchar(200)" json:"qqchat"`
	Wechat string `gorm:"type:varchar(100)" json:"wechat"`
	Weibo  string `gorm:"type:varchar(200)" json:"weibo"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
}

func GetProfile(id uint8) (Profile, int) {
	var profile Profile
	err = db.
		Where("ID =?", id).
		First(&profile).
		Error
	if err != nil {
		//关于 nil 可表示的值，参见 https://gfw.go101.org/article/nil.html。
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}
func UpdateProfile(id uint8, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
