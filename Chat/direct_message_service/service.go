package messageservice

import (
	"time"

	"github.com/google/uuid"
)

type DirectMessageService interface {
	// Returns all Messages that a readable for a user (userID)
	GetMessages(userID uuid.UUID) ([]Message, error)
	SendMessage(msg Message) error
	ReplaceMessage(messageID uuid.UUID, msg Message) error
	DeleteMessage(messageID uuid.UUID) error
}

type Reaction struct {
	Sender   uuid.UUID `json:"sender"`
	Reaction string    `json:"reaction"`
}

type Message struct {
	Id         uuid.UUID  `json:"id"`
	Content    string     `json:"content"`
	SenderID   uuid.UUID  `json:"sender"`
	ReceiverID uuid.UUID  `json:"receiver"`
	Timestamp  time.Time  `json:"timestamp"`
	ImageUrl   string     `json:"image"`
	Reactions  []Reaction `json:"reactions"`
}
