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
	Id       ID
	Title    string
	ImageUrl string
	Members  []Member
}

type Message struct {
	Id        ID
	Content   string
	Sender    Member
	Timestamp time.Time
	imageUrl  string
}

type Member struct {
	Id       ID
	name     string
	imageUrl string
}
