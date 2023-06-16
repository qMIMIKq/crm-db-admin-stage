package services

import (
	"crm_admin/internal/domain"
	"crm_admin/internal/repository"
	"crm_admin/pkg/hasher"
)

type AuthService struct {
	repo repository.Authorization
}

const salt = "asda312das"

func (a AuthService) CheckUser(user domain.UserAuth) (string, error) {
	user.Password = hasher.GeneratePasswordHash(user.Password, salt)
	return a.repo.CheckUser(user)
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
