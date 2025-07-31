package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	UserLogin
	Name      string `json:"name"`
	Phone     string `json:"phone"`
}
