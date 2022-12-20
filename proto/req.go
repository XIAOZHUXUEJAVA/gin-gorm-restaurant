package proto

type (
	ReqAddUser struct {
		// required： 必须填写, alpha： 必须是字母, min=5: 最小长度是4，max=12 最大长度是12, label可以翻译为用户名
		UserName string `json:"user_name" validate:"required,alpha,min=4,max=12" label:"用户名"`
		Email    string `json:"user_email" validate:"required,email" label:"用户邮箱"`
		Phone    string `json:"user_phone" validate:"required,customPhoneNumber" label:"用户电话"`
		Password string `json:"user_password"`
		Birth    string `json:"user_birth"`
		Gender   string `json:"user_gender"`
	}
	ReqAddTable struct {
		BookName   string `json:"book_name"`
		BookPhone  string `json:"book_phone"`
		BookPeople int    `json:"book_people"`
		BookTables int    `json:"book_tables"`
		UserId     int    `json:"user_id"`
		BookWhen   string `json:"book_when"`
		BookNote   string `json:"book_note"`
	}
)
