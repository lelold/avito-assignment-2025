package repository

import (
	"avito-assignment-2025/database"
	"avito-assignment-2025/model"
)

type IItemRepo interface {
	FetchByID(id uint) (*model.Item, error)
	FetchByName(name string) (*model.Item, error)
}

type ItemRepo struct{}

func NewItemRepo() IItemRepo {
	return &ItemRepo{}
}

func (r *ItemRepo) FetchByID(id uint) (*model.Item, error) {
	var item model.Item
	if err := database.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepo) FetchByName(name string) (*model.Item, error) {
	var item model.Item
	if err := database.DB.Where("name LIKE ?", "%"+name+"%").First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
