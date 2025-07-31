package service

import (
	"fmt"

	"github.com/coderemon24/go-ecommerce-app/internal/domain"
	"github.com/coderemon24/go-ecommerce-app/internal/dto"
	"github.com/coderemon24/go-ecommerce-app/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) FindUserByEmail(email string) (domain.User, error) {

	return domain.User{}, nil
}

func (s *UserService) Signup(input dto.UserRegister) (string, error) {

	user, err := s.Repo.UserCreate(domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", err
	}

	token := fmt.Sprintf("Bearer- %v,%v,%v", user.ID, user.Name, user.Email)

	return token, nil
}
