package service_test

import (
	"avito-assignment-2025/mocks"
	"avito-assignment-2025/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockTransactionRepo(t *testing.T) {
	mockRepo := new(mocks.MockTransactionRepo)
	transactions := []*model.Transaction{}
	mockRepo.On("FetchReceivedByID", uint(1)).Return(transactions, nil)

	res, err := mockRepo.FetchReceivedByID(1)
	assert.NoError(t, err)
	assert.Equal(t, transactions, res)
	mockRepo.AssertExpectations(t)
}
