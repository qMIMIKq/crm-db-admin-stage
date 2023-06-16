package services

import (
	"crm_admin/internal/domain"
	"crm_admin/internal/repository"
	"crm_admin/pkg/hasher"
)

type UsersServices struct {
	repo repository.Users
}

func (u UsersServices) EditUser(user domain.UserInfo) error {
	if len(user.Password) > 0 {
		user.Password = hasher.GeneratePasswordHash(user.Password, salt)
	}

	return u.repo.EditUser(user)
}

func (u UsersServices) CreateUser(user domain.UserInfo) (int, error) {
	user.Password = hasher.GeneratePasswordHash(user.Password, salt)
	return u.repo.CreateUser(user)
}

func (u UsersServices) GetUserByID(userId int) (domain.UserInfo, error) {
	return u.repo.GetUserByID(userId)
}

func (u UsersServices) GetUsers() (domain.Users, error) {
	return u.repo.GetUsers()
}

func NewUsersServices(repo repository.Users) *UsersServices {
	return &UsersServices{repo: repo}
}
