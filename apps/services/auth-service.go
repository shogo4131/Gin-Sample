package services

import (
	"go-gin-sample/apps/model"
	"go-gin-sample/apps/repository"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, paddword string) error
}

type AuthService struct {
	repository repository.IAuthRepository
}

func NewAuthService(repository repository.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := model.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repository.CreateUser(user)
}
