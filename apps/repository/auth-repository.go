package repository

import (
	"go-gin-sample/apps/model"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user model.User) error
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
