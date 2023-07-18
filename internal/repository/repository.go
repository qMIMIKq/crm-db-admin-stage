package repository

import (
	"crm_admin/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CheckUser(user domain.UserAuth) (string, error)
	CreateUser(user domain.UserCreate) (int, error)
}

type Users interface {
	CreateUser(user domain.UserInfo) (int, error)
	GetUsers() (domain.Users, error)
	GetUserByID(userId int) (domain.UserInfo, error)
	EditUser(user domain.UserInfo) error
}

type Groups interface {
	GetGroups() ([]domain.Group, error)
	GetGroupByID(groupID int) (domain.Group, error)
	EditGroup(group domain.Group) error
}

type Plots interface {
	GetPlots() ([]domain.Plot, error)
	GetPlotByID(plotId int) (domain.Plot, error)
	CreatePlot(plot domain.Plot) (int, error)
	EditPlot(plot domain.Plot) error
}

type Filters interface {
	GetFilters() ([]domain.FilterInfo, error)
	GetFilterByID(filterId int) (domain.FilterInfo, error)
	CreateFilter(filter domain.FilterInfo) (int, error)
	EditFilter(filter domain.FilterInfo) error
	UpdatePosition(filters []domain.FilterInfo) error
}

type Clients interface {
	GetClients() ([]domain.Client, error)
	GetClientByID(clientID int) (domain.Client, error)
	CreateClient(client domain.Client) (int, error)
	UpdateClient(client domain.Client) error
}

type Repository struct {
	Authorization
	Users
	Groups
	Plots
	Filters
	Clients
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPG(db),
		Users:         NewUsersPG(db),
		Groups:        NewGroupsPG(db),
		Plots:         NewPlotsPG(db),
		Filters:       NewFiltersPG(db),
		Clients:       NewClientsPG(db),
	}
}
