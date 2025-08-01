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

//	find user by email
func (s *UserService) FindUserByEmail(email string) (domain.User, error) {

	user, err :=s.Repo.FindUserByEmail(email)
	
	if(err != nil) {
		return domain.User{}, err
	}
	
	return user, nil
}

//	user signup -> create user
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

//	 find user by id
func (s *UserService) FindUserById(id uint) (domain.User, error) {
	user, err := s.Repo.GetUserById(id)
	
	if err != nil {
		return domain.User{}, err
	}
	
	return user, nil
}

//	find users
func (s *UserService) FindUsers() ([]domain.User, error) {
	users, err := s.Repo.GetAllUser()
	
	if err != nil {
		return []domain.User{}, err
	}
	
	return users, nil
}
