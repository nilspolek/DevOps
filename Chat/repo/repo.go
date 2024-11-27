package repo

import (
	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
)

type Repo interface {
	// Direct Messages
	//
	GetDirectMessages(userID uuid.UUID) ([]messageservice.Message, error)
	SendDirectMessage(msg messageservice.Message) error
	ReplaceDirecMessage(messageID uuid.UUID, msg messageservice.Message) error
	DeleteDirectMessage(uuid.UUID) error

	// Group Messages
	//
	GetGroupMessages(groupID uuid.UUID) ([]groupmessageservice.Message, error)
	SendMessageToGroup(groupID uuid.UUID, msg groupmessageservice.Message) error
	ReplaceGroupMessage(messageID uuid.UUID, msg groupmessageservice.Message) error
	DeleteGroupMessage(messageID uuid.UUID) error

	// Groups
	//
	GetAllGroups(userId uuid.UUID) ([]groupservice.Group, error)
	CreateGroup(group groupservice.Group) (uuid.UUID, error)
	EditGroup(group groupservice.Group, groupId uuid.UUID) error
	DeleteGroup(groupID uuid.UUID) error
	AddUserToGroup(groupId, userID uuid.UUID) error
	RemoveUserFromGroup(groupId, userID uuid.UUID) error

	// Reactions
	//
	AddReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction) error
	ChangeReactionToDM(messageID, userID uuid.UUID, reaction messageservice.Reaction) error
	RemoveReactionFromDM(messageID, userID uuid.UUID) error

	AddReactionToGroup(messageID, userID uuid.UUID, reaction groupmessageservice.Reaction) error
	ChangeReactionToGroup(messageID uuid.UUID, reaction groupmessageservice.Reaction) error
	RemoveReactionFromGroup(messageID, userID uuid.UUID) error
}
