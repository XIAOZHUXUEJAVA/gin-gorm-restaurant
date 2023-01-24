package model

import (
	"log"
	"zhu/myrest/proto"
	"zhu/myrest/utils/errmsg"
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

// 根据email获取user
func GetUserByEmail(email string) (*proto.RspFindUserByEmail, int) {
	var user *User = &User{}
	// db.First(&user, email).Select("user_id, user_name, user_password")
	err := db.Model(&user).Where("user_email = ?", email).First(&user).Select("user_id, user_name, user_password").Error
	if err != nil {
		log.Println(err)
		return nil, errmsg.ERROR
	}
	userInfo := &proto.RspFindUserByEmail{
		ID:       user.UserId,
		UserName: user.UserName,
		Password: user.UserPassword,
	}
	return userInfo, errmsg.SUCCESS
}

// 新增用户, 这个地方可以考虑优化一下，就是插入相同的话，返回错误信息，怎么办呢
func InsertUser(data *User) (*User, int) {
	// 判断用户是否存在
	userInfo, _ := GetUserByEmail(data.UserEmail)
	if userInfo != nil {
		return nil, errmsg.USER_EMAIL_EXIST
	}
	err := db.Create(&data).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS
}
