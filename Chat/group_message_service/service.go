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
	Id        ID         `json:"id"`
	Content   string     `json:"content"`
	Sender    Member     `json:"sender"`
	Timestamp time.Time  `json:"timestamp"`
	ImageUrl  string     `json:"image"`
	GroupId   ID         `json:"group_id"`
	Reactions []Reaction `json:"reactions"`
}

type Reaction struct {
	Sender   Member `json:"sender"`
	Reaction string `json:"reaction"`
}

type Member struct {
	Id       ID     `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image"`
}
