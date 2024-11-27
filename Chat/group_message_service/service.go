package groupmessageservice

import (
	"time"

	"github.com/google/uuid"
)

type GroupMessageService interface {
	// Returns all messages from a Group
	GetMessages(groupID uuid.UUID) ([]Message, error)
	SendMessage(groupID uuid.UUID, msg Message) error
	ReplaceMessage(messageID uuid.UUID, msg Message) error
	DeleteMessage(messageID uuid.UUID) error
}

type Message struct {
	Id        uuid.UUID  `json:"id"`
	Content   string     `json:"content"`
	Sender    Member     `json:"sender"`
	Timestamp time.Time  `json:"timestamp"`
	ImageUrl  string     `json:"image"`
	GroupId   uuid.UUID  `json:"group_id"`
	Reactions []Reaction `json:"reactions"`
}

type Reaction struct {
	Sender   Member `json:"sender"`
	Reaction string `json:"reaction"`
}

type Member struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ImageUrl string    `json:"image"`
}
