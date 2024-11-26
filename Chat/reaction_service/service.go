package reactionservice

import (
	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
)

type ReactionService interface {
	AddReactionToDM(messageID ID, reaction messageservice.Reaction) error
	ChangeReactionToDM(messageID ID, reaction messageservice.Reaction) error
	RemoveReactionFromDM(messageID, userID ID) error

	AddReactionToGroup(messageID, userId ID, reaction groupmessageservice.Reaction) error
	ChangeReactionToGroup(messageID ID, reaction groupmessageservice.Reaction) error
	RemoveReactionFromGroup(messageID, userID ID) error
}
type ID uuid.UUID
