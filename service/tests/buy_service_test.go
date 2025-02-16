package service_test

import (
	"avito-assignment-2025/mocks"
	"avito-assignment-2025/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockBuyRepo(t *testing.T) {
	mockRepo := new(mocks.MockBuyRepo)
	buy := &model.Buy{}
	mockRepo.On("Add", buy).Return(nil)

	err := mockRepo.Add(buy)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
