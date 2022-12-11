package model

import (
	"log"
	"zhu/myrest/errmsg"
)

type Food struct {
	FoodId       string `gorm:"type:int(11);" json:"food_id"  label:"菜品编号"`
	FoodName     string `gorm:"type:varchar(255);" json:"food_name"  label:"菜品名称"`
	FoodStar     string `gorm:"type:varchar(255);" json:"food_star"  label:"菜品评分"`
	FoodVote     string `gorm:"type:varchar(255);" json:"food_vote"  label:"菜品销量"`
	FoodPrice    string `gorm:"type:varchar(255);" json:"food_price"  label:"菜品价格"`
	FoodDiscount string `gorm:"type:varchar(255);" json:"food_discount"  label:"菜品折扣"`
	FoodDesc     string `gorm:"type:varchar(255);" json:"food_desc"  label:"菜品描述"`
	FoodStatus   string `gorm:"type:varchar(255);" json:"food_status"  label:"菜品状态"`
	FoodType     string `gorm:"type:varchar(255);" json:"food_type"  label:"菜品类型"`
	FoodCategory string `gorm:"type:varchar(255);" json:"food_category"  label:"菜品种类"`
	FoodSrc      string `gorm:"type:varchar(255);" json:"food_src"  label:"菜品图片"`
}

//	func GetAllFood() (foods []*Food, err error) {
//		if err = db.Find(&foods).Error; err != nil {
//			return nil, err
//		}
//		return
//	}
func GetAllFood() ([]*Food, int) {
	var foods []*Food
	// 如果查找失败，即err不为空
	err := db.Find(&foods).Error
	if err != nil {
		log.Println(err)
		return nil, errmsg.ERROR
	}
	return foods, errmsg.SUCCESS
}

func GetFoodById(id int) (*Food, int) {
	var food *Food
	err := db.Model(&food).Where("id = ?", id).First(&food).Error
	if err != nil {
		log.Println(err)
		return food, errmsg.ERROR
	}
	return food, errmsg.SUCCESS
}
