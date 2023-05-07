package model

import (
	"log"
	"zhu/myrest/proto"
	"zhu/myrest/utils/errmsg"

	"github.com/ulule/deepcopier"
)

type User struct {
	// validate:"required,min=4,max=12"
	UserId       int    `gorm:"type:int(11) auto_increment; primary_key; comment: '用户编号'" json:"user_id"  `
	UserName     string `gorm:"type:varchar(255); comment: '用户名';" json:"user_name"  `
	UserEmail    string `gorm:"type:varchar(255); comment: '用户邮箱';" json:"user_email"  `
	UserPhone    string `gorm:"type:varchar(255); comment: '用户电话';" json:"user_phone"  `
	UserPassword string `gorm:"type:varchar(255); comment: '用户密码';" json:"user_password"  `
	UserBirth    string `gorm:"type:varchar(255); comment: '用户生日';" json:"user_birth"  `
	UserGender   string `gorm:"type:varchar(255); comment: '用户性别';" json:"user_gender"  `
}

// 获取所有user
func GetAllUsers() (users []*User, err error) {
	if err = db.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

// 根据email获取userInfo
func GetUserByEmail(email string) (*proto.RspFindUserByEmail, int) {
	user := &User{}

	err := db.Model(user).
		Select("user_id, user_name, user_password").
		Where("user_email = ?", email).
		First(user).
		Error
	if err != nil {
		log.Printf("根据邮箱查找用户时出错: %s", err)
		return nil, errmsg.ERROR
	}

	userInfo := &proto.RspFindUserByEmail{}
	if err := deepcopier.Copy(user).To(userInfo); err != nil {
		log.Printf("转换用户信息时出错: %s", err)
		return nil, errmsg.ERROR
	}
	return userInfo, errmsg.SUCCESS
}

func InsertUser(data *User) (*User, int) {
	// 判断用户是否存在
	userInfo, _ := GetUserByEmail(data.UserEmail)
	if userInfo != nil {
		log.Println("用户邮箱已存在")
		return nil, errmsg.USER_EMAIL_EXIST
	}
	if err := db.Create(&data).Error; err != nil {
		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS
}
