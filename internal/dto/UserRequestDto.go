package dto

type UserLogin struct {
	Email    string	`json:"email"`
	Password string	`json:"password"`
}

type UserRegister struct {
	UserLogin
	Name    string	`json:"name"`
	Role    string	`json:"role"`
	Address string	`json:"address"`
	Phone   string	`json:"phone"`
}
