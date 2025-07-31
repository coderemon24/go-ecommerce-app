package dto

type UserLogin struct {
	Email    string
	Password string
}

type UserRegister struct {
	UserLogin
	Name    string
	Role    string
	Address string
	Phone   string
}
