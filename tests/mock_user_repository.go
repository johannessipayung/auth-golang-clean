package tests

import (
	"auth-golang-clean/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	args := m.Called(email)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id uint) (*entity.User, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}
