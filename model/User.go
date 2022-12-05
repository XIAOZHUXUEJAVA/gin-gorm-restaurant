package model

type User struct {
	// validate:"required,min=4,max=12"
	UserId       string `gorm:"primaryKey;type:int(11);not null;" json:"userId"  label:"用户编号"`
	UserName     string `gorm:"type:varchar(255);not null" json:"username"  label:"用户名"`
	UserEmail    string `gorm:"type:varchar(255);not null" json:"email"  label:"用户邮箱"`
	UserPhone    string `gorm:"type:varchar(255);not null" json:"phone"  label:"用户电话"`
	UserPassword string `gorm:"type:varchar(255);not null" json:"password"  label:"用户密码"`
	UserBirth    string `gorm:"type:varchar(255);not null" json:"birth"  label:"用户生日"`
	UserGender   string `gorm:"type:varchar(255);not null" json:"gender"  label:"用户性别"`
}

// 获取所有user
func GetAllUsers() (users []*User, err error) {
	if err = db.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}
