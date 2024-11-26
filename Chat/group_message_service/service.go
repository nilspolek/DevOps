package groupmessageservice

import (
	"time"

	"github.com/google/uuid"
)

type GroupMessageService interface {
	// Returns all messages from a Group
	GetMessages(groupID ID) ([]Message, error)
	SendMessage(groupID ID, msg Message) error
	ReplaceMessage(messageID ID, msg Message) error
	DeleteMessage(messageID ID) error
}

type ID uuid.UUID

type Message struct {
	Id        ID
	Content   string
	Sender    Member
	Timestamp time.Time
	imageUrl  string
	GroupId  ID
	Reactions []Reaction
}

type Reaction struct {
	Sender Member
	Reaction string
}

type Member struct {
	Id       ID
	name     string
	imageUrl string
}
