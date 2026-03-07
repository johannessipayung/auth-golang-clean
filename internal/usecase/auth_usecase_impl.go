package usecase

import (
	"auth-golang-clean/internal/entity"
	"auth-golang-clean/internal/repository"
	"auth-golang-clean/internal/utils"
	"errors"
)

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {

	return &authUsecase{repo}
}

func (u *authUsecase) Register(username, email, password string) error {

	hash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	user := entity.User{

		Username: username,
		Email:    email,
		Password: hash,
		Role:     "user",
	}

	return u.userRepo.Create(&user)
}

func (u *authUsecase) Login(email, password string) (string, string, error) {

	user, err := u.userRepo.FindByEmail(email)

	if err != nil {
		return "", "", err
	}

	err = utils.CheckPassword(user.Password, password)

	if err != nil {

		return "", "", errors.New("invalid credentials")
	}

	token, _ := utils.GenerateToken(user.Username, user.Email, user.Role)

	refresh, _ := utils.GenerateRefreshToken(user.Email)

	return token, refresh, nil
}
