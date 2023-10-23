package model

type User struct {
	Id       int32  `json:"id"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
