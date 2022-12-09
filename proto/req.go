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
)
