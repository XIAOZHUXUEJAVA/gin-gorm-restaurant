package proto

type (
	ReqAddUser struct {
		UserName string `json:"user_name"`
		Email    string `json:"user_email"`
		Phone    string `json:"user_phone"`
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
