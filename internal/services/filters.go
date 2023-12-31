package services

import (
	"crm_admin/internal/domain"
	"crm_admin/internal/repository"
)

type FiltersService struct {
	repo repository.Filters
}

func (f FiltersService) EditFilter(filter domain.FilterInfo) error {
	return f.repo.EditFilter(filter)
}

func (f FiltersService) CreateFilter(filter domain.FilterInfo) (int, error) {
	return f.repo.CreateFilter(filter)
}

func (f FiltersService) GetFilters() ([]domain.FilterInfo, error) {
	return f.repo.GetFilters()
}

func (f FiltersService) GetFilterByID(filterId int) (domain.FilterInfo, error) {
	return f.repo.GetFilterByID(filterId)
}

func NewFiltersService(repo repository.Filters) *FiltersService {
	return &FiltersService{repo: repo}
}
