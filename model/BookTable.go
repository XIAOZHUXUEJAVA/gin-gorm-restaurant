package model

import "zhu/myrest/errmsg"

type BookTable struct {
	BookId     int    `gorm:"type:int(11) auto_increment; comment: '预订编号';  primary_key;" json:"book_id"  `
	BookName   string `gorm:"type:varchar(255); comment: '姓名';" json:"book_name"  `
	BookPhone  string `gorm:"type:varchar(255); comment: '预订电话';" json:"book_phone"  `
	BookPeople int    `gorm:"type:int(11); comment: '预订人数';" json:"book_people"  `
	BookTables int    `gorm:"type:int(11); comment: '预订桌数';" json:"book_tables"  `
	UserId     int    `gorm:"type:int(11); comment: '用户编号';" json:"user_id"  `
	BookWhen   string `gorm:"type:varchar(255); comment: '预订时间';" json:"book_when"  `
	BookNote   string `gorm:"type:text; comment: '备注';" json:"book_note"  `
}

func InsertTable(data *BookTable) (*BookTable, int) {
	err := db.Create(&data).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS

}
