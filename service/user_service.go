package service

import (
	"avito-assignment-2025/middleware"
	"avito-assignment-2025/model"
	"avito-assignment-2025/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Authenticate(name, password string) (string, error)
}

type UserService struct {
	userRepo repository.IUserRepo
}

func NewUserService(userRepo repository.IUserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Authenticate(name, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(name)
	if err != nil {
		hashedPassword, hashErr := hashPassword(password)
		if hashErr != nil {
			return "", hashErr
		}
		user = &model.User{
			Username: name,
			Password: hashedPassword,
			Balance:  1000,
		}
		if createErr := s.userRepo.Create(user); createErr != nil {
			return "", createErr
		}
	}

	if !comparePasswords(user.Password, password) {
		return "", errors.New("неверный пароль")
	}

	token, tokenErr := middleware.CreateToken(user.ID, user.Username)
	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func comparePasswords(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
