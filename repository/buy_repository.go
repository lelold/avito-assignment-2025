package repository

import (
	"avito-assignment-2025/database"
	"avito-assignment-2025/model"
	"fmt"
)

type IBuyRepo interface {
	Add(buy *model.Buy) error
	FetchByUserAndItem(userID uint, itemID uint) (*model.Buy, error)
	FetchListByUserID(userID uint) ([]*model.Buy, error)
	Update(buy model.Buy) error
}

type buyRepo struct{}

func NewBuyRepo() IBuyRepo {
	return &buyRepo{}
}

func (r *buyRepo) Add(buy *model.Buy) error {
	return database.DB.Create(buy).Error
}

func (r *buyRepo) FetchByUserAndItem(userID uint, itemID uint) (*model.Buy, error) {
	var buy model.Buy
	if err := database.DB.Where("user_id = ? AND item_id = ?", userID, itemID).First(&buy).Error; err != nil {
		return nil, err
	}
	return &buy, nil
}

func (r *buyRepo) FetchListByUserID(userID uint) ([]*model.Buy, error) {
	var buys []*model.Buy
	if err := database.DB.Where("user_id = ?", userID).Find(&buys).Error; err != nil {
		return nil, err
	}
	return buys, nil
}

func (r *buyRepo) Update(buy model.Buy) error {
	result := database.DB.Model(&model.Buy{}).
		Where("user_id = ? AND item_id = ?", buy.UserID, buy.ItemID).
		Update("count", buy.Count)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("покупка не найдена")
	}

	return nil
}
