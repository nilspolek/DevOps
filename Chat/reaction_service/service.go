package reactionservice

import "github.com/google/uuid"

type ReactionService interface {
	AddReaction(messageID ID, reaction string) error
	ChangeReaction(messageID ID, reaction string) error
	RemoveReaction(messageID ID) error
}
type ID uuid.UUID
