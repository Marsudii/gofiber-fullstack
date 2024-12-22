package service

import "gofiber-fullstack/internal/repository"

type AuthService interface {
	Login(email string, password string) (string, error)
}

type authService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(authRepo repository.AuthRepo) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) Login(email string, password string) (string, error) {
	return s.authRepo.Login(email, password)
}
