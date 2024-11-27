package gimpl

import (
	"github.com/google/uuid"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
	"github.com/nilspolek/DevOps/Chat/repo"
)

type svc struct {
	repo repo.Repo
}

func New(repo repo.Repo) groupservice.GroupService {
	return &svc{repo: repo}
}

func (s *svc) GetAllGroups(userId uuid.UUID) ([]groupservice.Group, error) {
	return s.repo.GetAllGroups(userId)
}

func (s *svc) CreateGroup(group groupservice.Group) (uuid.UUID, error) {
	return s.repo.CreateGroup(group)
}

func (s *svc) EditGroup(group groupservice.Group, id uuid.UUID) error {
	return s.repo.EditGroup(group, id)
}

func (s *svc) DeleteGroup(id uuid.UUID) error {
	return s.repo.DeleteGroup(id)
}

func (s svc) AddUserToGroup(groupId, userId uuid.UUID) error {
	return s.repo.AddUserToGroup(groupId, userId)
}

func (s svc) RemoveUserFromGroup(groupId, userId uuid.UUID) error {
	return s.repo.RemoveUserFromGroup(groupId, userId)
}
