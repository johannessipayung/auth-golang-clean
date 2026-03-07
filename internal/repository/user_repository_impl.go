package repository

import (
	"auth-golang-clean/internal/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &userRepository{db}
}

func (r *userRepository) Create(user *entity.User) error {

	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {

	var user entity.User

	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}
