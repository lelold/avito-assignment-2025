package service

import (
	"avito-assignment-2025/database"
	"avito-assignment-2025/model"
	"avito-assignment-2025/repository"
	"fmt"
)

type IBuyService interface {
	BuyItem(userName string, itemName string) error
}

type buyService struct {
	userRepo repository.IUserRepo
	itemRepo repository.IItemRepo
	buyRepo  repository.IBuyRepo
}

func NewBuyService(
	userRepo repository.IUserRepo,
	itemRepo repository.IItemRepo,
	buyRepo repository.IBuyRepo,
) *buyService {
	return &buyService{
		userRepo: userRepo,
		itemRepo: itemRepo,
		buyRepo:  buyRepo,
	}
}

func (s *buyService) BuyItem(userName string, itemName string) error {
	user, err := s.userRepo.FindByUsername(userName)
	if err != nil {
		return fmt.Errorf("пользователь не найден")
	}

	item, err := s.itemRepo.FetchByName(itemName)
	if err != nil {
		return fmt.Errorf("предмет не найден")
	}

	if user.Balance < item.Price {
		return fmt.Errorf("недостаточно средств")
	}

	tx := database.DB.Begin()
	user.Balance -= item.Price

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	buy, err := s.buyRepo.FetchByUserAndItem(user.ID, item.ID)
	if err != nil {
		buy = &model.Buy{
			UserID: user.ID,
			ItemID: item.ID,
			Count:  1,
		}
		if err := s.buyRepo.Add(buy); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		buy.Count++
		if err := s.buyRepo.Update(*buy); err != nil {
			tx.Rollback()
			return fmt.Errorf("невозможно обновить данные")
		}
	}

	return tx.Commit().Error
}
