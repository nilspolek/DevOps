package groupservice

import (
	"time"

	"github.com/google/uuid"
)

type GroupService interface {
	// Returns all Groups an user has access to
	GetAllGroups(userID ID) ([]Group, error)
	CreateGroup(group Group) (ID, error)
	EditGroup(group Group, id ID) error
	DeleteGroup(id ID) error
	AddUserToGroup(groupId, userId ID) error
	RemoveUserFromGroup(groupId, userId ID) error
}

type ID uuid.UUID

type Group struct {
	Id       ID       `json:"id"`
	Title    string   `json:"title"`
	ImageUrl string   `json:"image"`
	Members  []Member `json:"members"`
}

type Message struct {
	Id        ID        `json:"id"`
	Content   string    `json:"content"`
	Sender    Member    `json:"sender"`
	Timestamp time.Time `json:"timestamp"`
	imageUrl  string    `json:"image"`
}

type Member struct {
	Id       ID     `json:"id"`
	name     string `json:"name"`
	imageUrl string `json:"image"`
}
