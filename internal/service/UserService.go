package service

import (
	"log"

	"github.com/coderemon24/go-ecommerce-app/internal/domain"
	"github.com/coderemon24/go-ecommerce-app/internal/dto"
	"gorm.io/gorm"
)

type UserService struct{
	DB *gorm.DB
}

func (s *UserService) FindUserByEmail(email string) (domain.User, error) {
	
	return domain.User{}, nil
}

func (s *UserService) Signup(input dto.UserRegister)(string, error){
	
	log.Println(input)
	
	token := "user-register-token";
	
	return token, nil
}