package mocks

import (
	"avito-assignment-2025/model"

	"github.com/stretchr/testify/mock"
)

type MockBuyRepo struct {
	mock.Mock
}

func (m *MockBuyRepo) Add(buy *model.Buy) error {
	args := m.Called(buy)
	return args.Error(0)
}

func (m *MockBuyRepo) FetchByUserAndItem(userID uint, itemID uint) (*model.Buy, error) {
	args := m.Called(userID, itemID)
	return args.Get(0).(*model.Buy), args.Error(1)
}

func (m *MockBuyRepo) FetchListByUserID(userID uint) ([]*model.Buy, error) {
	args := m.Called(userID)
	return args.Get(0).([]*model.Buy), args.Error(1)
}

func (m *MockBuyRepo) Update(buy model.Buy) error {
	args := m.Called(buy)
	return args.Error(0)
}

type MockItemRepo struct {
	mock.Mock
}

func (m *MockItemRepo) FetchByID(id uint) (*model.Item, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Item), args.Error(1)
}

func (m *MockItemRepo) FetchByName(name string) (*model.Item, error) {
	args := m.Called(name)
	return args.Get(0).(*model.Item), args.Error(1)
}

type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) Add(tx *model.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockTransactionRepo) FetchReceivedByID(id uint) ([]*model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).([]*model.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) FetchSentByID(id uint) ([]*model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).([]*model.Transaction), args.Error(1)
}

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) FindByID(ID uint) (*model.User, error) {
	args := m.Called(ID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepo) FindByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}
