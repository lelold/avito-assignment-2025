package service

import (
	"avito-assignment-2025/database"
	"avito-assignment-2025/model"
	"avito-assignment-2025/repository"
	"fmt"
)

type ITransactionService interface {
	TransferCoins(senderName, receiverName string, amount int) error
}

type transactionService struct {
	userRepo        repository.IUserRepo
	transactionRepo repository.ITransactionRepo
}

func NewTransactionService(
	userRepo repository.IUserRepo,
	txRepo repository.ITransactionRepo,
) *transactionService {
	return &transactionService{
		userRepo:        userRepo,
		transactionRepo: txRepo,
	}
}

func (s *transactionService) TransferCoins(senderName, receiverName string, amount int) error {
	if senderName == receiverName {
		return fmt.Errorf("нельзя перевести монеты самому себе")
	}

	sender, err := s.userRepo.FindByUsername(senderName)
	if err != nil {
		return fmt.Errorf("отправитель не найден")
	}

	receiver, err := s.userRepo.FindByUsername(receiverName)
	if err != nil {
		return fmt.Errorf("получатель не найден")
	}

	if sender.Balance < amount {
		return fmt.Errorf("недостаточно средств")
	}

	transact := database.DB.Begin()

	sender.Balance -= amount
	receiver.Balance += amount

	if err := transact.Save(sender).Error; err != nil {
		transact.Rollback()
		return err
	}

	if err := transact.Save(receiver).Error; err != nil {
		transact.Rollback()
		return err
	}

	transaction := model.Transaction{
		FromUser: sender.ID,
		ToUser:   receiver.ID,
		Amount:   amount,
	}

	if err := transact.Create(&transaction).Error; err != nil {
		transact.Rollback()
		return err
	}

	return transact.Commit().Error
}
