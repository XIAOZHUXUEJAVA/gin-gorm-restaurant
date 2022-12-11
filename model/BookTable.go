package model

import "zhu/myrest/errmsg"

type BookTable struct {
	BookId     int    `gorm:"type:int(11);primary_key;AUTO_INCREMENT" json:"book_id"  label:"预订编号"`
	BookName   string `gorm:"type:varchar(255);" json:"book_name"  label:"姓名"`
	BookPhone  string `gorm:"type:varchar(255);" json:"book_phone"  label:"预订电话"`
	BookPeople int    `gorm:"type:int(11);" json:"book_people"  label:"预订人数"`
	BookTables int    `gorm:"type:int(11);" json:"book_tables"  label:"预订桌数"`
	UserId     int    `gorm:"type:int(11);" json:"user_id"  label:"用户编号"`
	BookWhen   string `gorm:"type:varchar(255);" json:"book_when"  label:"预订时间"`
	BookNote   string `gorm:"type:text;" json:"book_note"  label:"备注"`
}

func InsertTable(data *BookTable) (*BookTable, int) {
	err := db.Create(&data).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS

}
