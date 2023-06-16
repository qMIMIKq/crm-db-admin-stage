package services

import (
	"crm_admin/internal/domain"
	"crm_admin/internal/repository"
)

type ClientsServices struct {
	repo repository.Clients
}

func (c ClientsServices) GetClients() ([]domain.Client, error) {
	return c.repo.GetClients()
}

func (c ClientsServices) GetClientByID(clientID int) (domain.Client, error) {
	return c.repo.GetClientByID(clientID)
}

func (c ClientsServices) CreateClient(client domain.Client) (int, error) {
	return c.repo.CreateClient(client)
}

func (c ClientsServices) UpdateClient(client domain.Client) error {
	return c.repo.UpdateClient(client)
}

func NewClientsServices(repo repository.Clients) *ClientsServices {
	return &ClientsServices{repo: repo}
}
