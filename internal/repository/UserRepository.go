package repository

import (
	"errors"

	"github.com/coderemon24/go-ecommerce-app/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(email string) (domain.User, error)
	UserCreate(usr domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) UserCreate(usr domain.User) (domain.User, error) {
	
	err := r.db.Create(&usr).Error
	
	if err != nil {
		return domain.User{}, errors.New("failed to create user")
	}
	
	return usr, err
}

func (r *userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	
	err := r.db.Where("email = ?", email).First(&user).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, errors.New("user not found")
		}
		
		return domain.User{}, errors.New("failed to find user")
	}
	
	return user, err
}