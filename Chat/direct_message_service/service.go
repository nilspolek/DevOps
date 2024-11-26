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
	Sender   ID
	Reaction string
}

type Message struct {
	Id         ID
	Content    string
	SenderID   ID
	ReceiverID ID
	Timestamp  time.Time
	ImageUrl   string
	Reactions  []Reaction
}
