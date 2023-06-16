package services

import (
	"crm_admin/internal/domain"
	"crm_admin/internal/repository"
)

type GroupsService struct {
	repo repository.Groups
}

func (g GroupsService) EditGroup(group domain.Group) error {
	return g.repo.EditGroup(group)
}

func (g GroupsService) GetGroupByID(groupID int) (domain.Group, error) {
	return g.repo.GetGroupByID(groupID)
}

func (g GroupsService) GetGroup(groupName string) {
	//TODO implement me
	panic("implement me")
}

func (g GroupsService) GetGroups() ([]domain.Group, error) {
	return g.repo.GetGroups()
}

func NewGroupsService(repo repository.Groups) *GroupsService {
	return &GroupsService{repo: repo}
}
