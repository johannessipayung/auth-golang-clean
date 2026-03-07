package tests

import (
	"auth-golang-clean/internal/entity"
	"auth-golang-clean/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepo struct{}

func (m *MockRepo) Create(user *entity.User) error {

	return nil
}

func (m *MockRepo) FindByEmail(email string) (*entity.User, error) {

	return &entity.User{

		Username: "test",

		Email: email,

		Password: "$2a$14$wHhR8V3g9yK/0Q1Z4eU9ueHfXx3O8DqRLVQjaxAg/P070MqxsVXni",

		Role: "user",
	}, nil
}

func TestRegister(t *testing.T) {

	mockRepo := new(MockRepo)

	auth := usecase.NewAuthUsecase(mockRepo)

	err := auth.Register("test", "test@gmail.com", "123456")

	assert.Nil(t, err)
}
