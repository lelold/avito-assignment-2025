package service_test

import (
	"avito-assignment-2025/mocks"
	"avito-assignment-2025/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockUserRepo(t *testing.T) {
	mockRepo := new(mocks.MockUserRepo)
	user := &model.User{}
	mockRepo.On("FindByID", uint(1)).Return(user, nil)

	res, err := mockRepo.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, res)
	mockRepo.AssertExpectations(t)
}
