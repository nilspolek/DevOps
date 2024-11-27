package groupservice

import (
	"time"

	"github.com/google/uuid"
)

type GroupService interface {
	// Returns all Groups an user has access to
	GetAllGroups(userID, authUser uuid.UUID) ([]Group, error)
	CreateGroup(group Group, authUser uuid.UUID) (uuid.UUID, error)
	EditGroup(group Group, id, authUser uuid.UUID) error
	DeleteGroup(id, authUser uuid.UUID) error
	AddUserToGroup(groupId, userId, authUser uuid.UUID) error
	RemoveUserFromGroup(groupId, userId, authUser uuid.UUID) error
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
