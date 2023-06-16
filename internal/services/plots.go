package services

import (
	"crm_admin/internal/domain"
	"crm_admin/internal/repository"
)

type PlotsService struct {
	repo repository.Plots
}

func (p PlotsService) EditPlot(plot domain.Plot) error {
	return p.repo.EditPlot(plot)
}

func (p PlotsService) GetPlotByID(plotId int) (domain.Plot, error) {
	return p.repo.GetPlotByID(plotId)
}

func (p PlotsService) CreatePlot(plot domain.Plot) (int, error) {
	return p.repo.CreatePlot(plot)
}

func (p PlotsService) GetPlots() ([]domain.Plot, error) {
	return p.repo.GetPlots()
}

func NewPlotsService(repo repository.Plots) *PlotsService {
	return &PlotsService{repo: repo}
}
