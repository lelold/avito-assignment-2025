package service

import (
	"avito-assignment-2025/model"
	"avito-assignment-2025/repository"
	"fmt"
)

type IInfoService interface {
	GetInfo(userID uint) (*model.InfoResponse, error)
}

type infoService struct {
	userRepo repository.IUserRepo
	txRepo   repository.ITransactionRepo
	itemRepo repository.IItemRepo
	buyRepo  repository.IBuyRepo
}

func NewInfoService(
	userRepo repository.IUserRepo,
	txRepo repository.ITransactionRepo,
	itemRepo repository.IItemRepo,
	buyRepo repository.IBuyRepo,
) *infoService {
	return &infoService{
		userRepo: userRepo,
		txRepo:   txRepo,
		itemRepo: itemRepo,
		buyRepo:  buyRepo,
	}
}

func (s *infoService) GetInfo(userID uint) (*model.InfoResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден")
	}

	var info model.InfoResponse
	info.Coins = user.Balance

	inventory, _ := s.buyRepo.FetchListByUserID(user.ID)
	info.Inventory = make([]model.InventoryItem, 0, len(inventory))
	for _, p := range inventory {
		item, err := s.itemRepo.FetchByID(p.ItemID)
		if err != nil {
			continue
		}
		info.Inventory = append(info.Inventory, model.InventoryItem{
			Type:     item.Name,
			Quantity: p.Count,
		})
	}

	received, _ := s.txRepo.FetchReceivedByID(user.ID)
	info.CoinHistory.Received = make([]model.CoinTransaction, 0, len(received))
	for _, t := range received {
		fromUser, err := s.userRepo.FindByID(t.FromUser)
		if err != nil {
			return nil, fmt.Errorf("не найден отправитель")
		}
		info.CoinHistory.Received = append(info.CoinHistory.Received, model.CoinTransaction{
			FromUser: fromUser.Username,
			Amount:   uint(t.Amount),
		})
	}

	sent, _ := s.txRepo.FetchSentByID(user.ID)
	info.CoinHistory.Sent = make([]model.CoinTransaction, 0, len(sent))
	for _, t := range sent {
		toUser, err := s.userRepo.FindByID(t.ToUser)
		if err != nil {
			return nil, fmt.Errorf("не найден получатель")
		}
		info.CoinHistory.Sent = append(info.CoinHistory.Sent, model.CoinTransaction{
			ToUser: toUser.Username,
			Amount: uint(t.Amount),
		})
	}

	return &info, nil
}
