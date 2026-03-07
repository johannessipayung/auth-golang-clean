package usecase

type AuthUsecase interface {

	Register(username, email, password string) error

	Login(email, password string) (string, string, error)
}
