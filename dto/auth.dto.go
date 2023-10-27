package dto

type AuthDto struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
