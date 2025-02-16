package repository

import (
	"avito-assignment-2025/database"
	"avito-assignment-2025/model"
)

type IUserRepo interface {
	Create(user *model.User) error
	FindByID(ID uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}

type userRepo struct{}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepo) FindByID(ID uint) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("id = ?", ID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
