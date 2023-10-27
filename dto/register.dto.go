package dto

type RegisterUser struct {
	Name      string `form:"name"`
	Password  string `form:"password"`
	Cpassword string `form:"c_password"`
	Email     string `form:"email"`
	Role      string `form:"role"`
}
