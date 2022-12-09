package model

import (
	"log"
	"zhu/myrest/errmsg"
	"zhu/myrest/proto"
)

type User struct {
	// validate:"required,min=4,max=12"
	UserId       int    `gorm:"primaryKey;type:int(11);not null;" json:"user_id"  label:"用户编号"`
	UserName     string `gorm:"type:varchar(255);not null" json:"user_name"  label:"用户名"`
	UserEmail    string `gorm:"type:varchar(255);not null" json:"user_email"  label:"用户邮箱"`
	UserPhone    string `gorm:"type:varchar(255);not null" json:"user_phone"  label:"用户电话"`
	UserPassword string `gorm:"type:varchar(255);not null" json:"user_password"  label:"用户密码"`
	UserBirth    string `gorm:"type:varchar(255);not null" json:"user_birth"  label:"用户生日"`
	UserGender   string `gorm:"type:varchar(255);not null" json:"user_gender"  label:"用户性别"`
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

// 新增用户
func InsertUser(data *User) (*User, int) {
	err := db.Create(&data).Error
	if err != nil {

		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS
}
