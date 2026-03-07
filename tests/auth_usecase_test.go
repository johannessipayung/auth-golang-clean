package tests

import (
	"auth-golang-clean/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {

	mockRepo := new(MockUserRepository)

	mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)

	authUsecase := usecase.NewAuthUsecase(mockRepo)

	err := authUsecase.Register(
		"testuser",
		"test@mail.com",
		"password",
	)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
