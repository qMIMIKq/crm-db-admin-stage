package services

import (
	"crm_admin/internal/domain"
	"crm_admin/internal/repository"
)

type Authorization interface {
	CheckUser(user domain.UserAuth) (string, error)
}

type Groups interface {
	GetGroup(groupName string)
	GetGroups() ([]domain.Group, error)
	GetGroupByID(groupID int) (domain.Group, error)
	EditGroup(group domain.Group) error
}

type Plots interface {
	GetPlotByID(plotId int) (domain.Plot, error)
	GetPlots() ([]domain.Plot, error)
	CreatePlot(plot domain.Plot) (int, error)
	EditPlot(plot domain.Plot) error
}

type Users interface {
	CreateUser(user domain.UserInfo) (int, error)
	GetUserByID(userId int) (domain.UserInfo, error)
	GetUsers() (domain.Users, error)
	EditUser(user domain.UserInfo) error
}

type Filters interface {
	GetFilters() ([]domain.FilterInfo, error)
	GetFilterByID(filterId int) (domain.FilterInfo, error)
	CreateFilter(filter domain.FilterInfo) (int, error)
	EditFilter(filter domain.FilterInfo) error
}

type Clients interface {
	GetClients() ([]domain.Client, error)
	GetClientByID(clientID int) (domain.Client, error)
	CreateClient(client domain.Client) (int, error)
	UpdateClient(client domain.Client) error
}

type Services struct {
	Authorization
	Users
	Groups
	Plots
	Filters
	Clients
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUsersServices(repos.Users),
		Groups:        NewGroupsService(repos.Groups),
		Plots:         NewPlotsService(repos.Plots),
		Filters:       NewFiltersService(repos.Filters),
		Clients:       NewClientsServices(repos.Clients),
	}
}
