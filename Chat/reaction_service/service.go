package reactionservice

import (
	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
)

type ReactionService interface {
	AddReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction, authUser uuid.UUID) error
	ChangeReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction, authUser uuid.UUID) error
	RemoveReactionFromDM(messageID, userID, authUser uuid.UUID) error

	AddReactionToGroup(messageID, userId uuid.UUID, reaction groupmessageservice.Reaction, authUser uuid.UUID) error
	ChangeReactionToGroup(messageID uuid.UUID, reaction groupmessageservice.Reaction, authUser uuid.UUID) error
	RemoveReactionFromGroup(messageID, userID, authUser uuid.UUID) error
}
