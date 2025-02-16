package service_test

import (
	"avito-assignment-2025/mocks"
	"avito-assignment-2025/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockItemRepo(t *testing.T) {
	mockRepo := new(mocks.MockItemRepo)
	item := &model.Item{}
	mockRepo.On("FetchByID", uint(1)).Return(item, nil)

	res, err := mockRepo.FetchByID(1)
	assert.NoError(t, err)
	assert.Equal(t, item, res)
	mockRepo.AssertExpectations(t)
}
