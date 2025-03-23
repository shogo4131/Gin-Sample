package repository

import (
	"errors"
	"go-gin-sample/apps/model"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user model.User) error
	FindUser(email string) (*model.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user model.User) error {
	result := r.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *AuthRepository) FindUser(email string) (*model.User, error) {
	var user model.User

	result := r.db.First(&user, "email = ?", email)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}
