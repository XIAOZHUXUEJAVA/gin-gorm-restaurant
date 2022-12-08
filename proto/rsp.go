package proto

type RspFindUserByEmail struct {
	ID       int    `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
}
