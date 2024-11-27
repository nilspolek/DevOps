package groupservice

import (
	"time"

	"github.com/google/uuid"
)

type GroupService interface {
	// Returns all Groups an user has access to
	GetAllGroups(userID uuid.UUID) ([]Group, error)
	CreateGroup(group Group) (uuid.UUID, error)
	EditGroup(group Group, id uuid.UUID) error
	DeleteGroup(id uuid.UUID) error
	AddUserToGroup(groupId, userId uuid.UUID) error
	RemoveUserFromGroup(groupId, userId uuid.UUID) error
}

type Group struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	ImageUrl string    `json:"image"`
	Members  []Member  `json:"members"`
}

type Message struct {
	Id        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Sender    Member    `json:"sender"`
	Timestamp time.Time `json:"timestamp"`
	ImageUrl  string    `json:"image"`
}

type Member struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ImageUrl string    `json:"image"`
}
