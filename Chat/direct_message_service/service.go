package messageservice

import (
	"time"

	"github.com/google/uuid"
)

type DirectMessageService interface {
	// Returns all Messages that a readable for a user (userID)
	GetMessages(userID ID) ([]Message, error)
	SendMessage(msg Message) error
	ReplaceMessage(messageID ID, msg Message) error
	DeleteMessage(messageID ID) error
}

type ID uuid.UUID

type Reaction struct {
	Sender   ID     `json:"sender"`
	Reaction string `json:"reaction"`
}

type Message struct {
	Id         ID         `json:"id"`
	Content    string     `json:"content"`
	SenderID   ID         `json:"sender"`
	ReceiverID ID         `json:"receiver"`
	Timestamp  time.Time  `json:"timestamp"`
	ImageUrl   string     `json:"image"`
	Reactions  []Reaction `json:"reactions"`
}
