package model

import (
	"log"
	"zhu/myrest/utils/errmsg"
)

type Food struct {
	FoodId       string `gorm:"type:int(11) auto_increment; primary_key;  comment: '菜品编号' " json:"food_id"  `
	FoodName     string `gorm:"type:varchar(255); comment: '菜品名称';" json:"food_name"`
	FoodStar     string `gorm:"type:varchar(255); comment: '菜品评分';" json:"food_star"  `
	FoodVote     string `gorm:"type:varchar(255); comment: '菜品销量';" json:"food_vote"  `
	FoodPrice    string `gorm:"type:varchar(255); comment: '菜品价格';" json:"food_price"  `
	FoodDiscount string `gorm:"type:varchar(255); comment: '菜品折扣';" json:"food_discount"  `
	FoodDesc     string `gorm:"type:varchar(255); comment: '菜品描述';" json:"food_desc"  `
	FoodStatus   string `gorm:"type:varchar(255); comment: '菜品状态';" json:"food_status"  `
	FoodType     string `gorm:"type:varchar(255); comment: '菜品类型';" json:"food_type"  `
	FoodCategory string `gorm:"type:varchar(255); comment: '菜品种类';" json:"food_category"  `
	FoodSrc      string `gorm:"type:varchar(255); comment: '菜品图片';" json:"food_src"  `
}

//	func GetAllFood() (foods []*Food, err error) {
//		if err = db.Find(&foods).Error; err != nil {
//			return nil, err
//		}
//		return
//	}
func GetAllFood() ([]*Food, int) {

	// var foods []*Food
	// // 如果查找失败，即err不为空
	// err := db.Find(&foods).Error
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, errmsg.ERROR
	// }
	// return foods, errmsg.SUCCESS

	var foods []*Food
	if err := db.Find(&foods).Error; err != nil {
		log.Printf("查找所有食物数据时出错: %s", err)
		return nil, errmsg.ERROR
	}
	return foods, errmsg.SUCCESS
}

func GetFoodById(id int) (*Food, int) {
	// var food *Food
	// err := db.Model(&food).Where("id = ?", id).First(&food).Error
	// if err != nil {
	// 	log.Println(err)
	// 	return food, errmsg.ERROR
	// }
	// return food, errmsg.SUCCESS
	var food *Food
	if err := db.Model(&food).Where("id = ?", id).First(&food).Error; err != nil {
		return nil, errmsg.ERROR
	}
	return food, errmsg.SUCCESS
}
