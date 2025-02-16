package repository

import (
	"avito-assignment-2025/database"
	"avito-assignment-2025/model"
)

type ITransactionRepo interface {
	Add(tx *model.Transaction) error
	FetchReceivedByID(id uint) ([]*model.Transaction, error)
	FetchSentByID(id uint) ([]*model.Transaction, error)
}

type transactionRepo struct{}

func NewTransactionRepo() ITransactionRepo {
	return &transactionRepo{}
}

func (r *transactionRepo) Add(tx *model.Transaction) error {
	return database.DB.Create(tx).Error
}

func (r *transactionRepo) FetchReceivedByID(id uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := database.DB.Where("to_user = ?", id).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}

func (r *transactionRepo) FetchSentByID(id uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := database.DB.Where("from_user = ?", id).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}
