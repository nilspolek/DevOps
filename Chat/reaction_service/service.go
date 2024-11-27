package reactionservice

import (
	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
)

type ReactionService interface {
	AddReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction) error
	ChangeReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction) error
	RemoveReactionFromDM(messageID, userID uuid.UUID) error

	AddReactionToGroup(messageID, userId uuid.UUID, reaction groupmessageservice.Reaction) error
	ChangeReactionToGroup(messageID uuid.UUID, reaction groupmessageservice.Reaction) error
	RemoveReactionFromGroup(messageID, userID uuid.UUID) error
}
