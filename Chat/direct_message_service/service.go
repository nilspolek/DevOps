package messageservice

import (
	"time"

	"github.com/google/uuid"
)

type DirectMessageService interface {
	// Returns all Messages that a readable for a user (userID)
	GetMessages(userID, authUser uuid.UUID) ([]Message, error)
	SendMessage(msg Message, authUser uuid.UUID) error
	ReplaceMessage(messageID uuid.UUID, msg Message, authUser uuid.UUID) error
	DeleteMessage(messageID uuid.UUID, authUser uuid.UUID) error
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
