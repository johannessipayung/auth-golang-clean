package repository

import "auth-golang-clean/internal/entity"

type UserRepository interface {

	Create(user *entity.User) error

	FindByEmail(email string) (*entity.User, error)
}
