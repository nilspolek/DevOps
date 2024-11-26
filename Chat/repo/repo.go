package repo

import (
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
)

type Repo interface {
	// Direct Messages
	//
	GetDirectMessages(userID messageservice.ID) ([]messageservice.Message, error)
	SendDirectMessage(msg messageservice.Message) error
	ReplaceDirecMessage(messageID messageservice.ID, msg messageservice.Message) error
	DeleteDirectMessage(messageservice.ID) error

	// Group Messages
	//
	GetGroupMessages(groupID groupmessageservice.ID) ([]groupmessageservice.Message, error)
	SendMessageToGroup(groupID groupmessageservice.ID, msg groupmessageservice.Message) error
	ReplaceGroupMessage(messageID groupmessageservice.ID, msg groupmessageservice.Message) error
	DeleteGroupMessage(messageID groupmessageservice.ID) error

	// Groups
	//
	GetAllGroups(userId groupservice.ID) ([]groupservice.Group, error)
	CreateGroup(group groupservice.Group) (groupservice.ID, error)
	EditGroup(group groupservice.Group, groupId groupservice.ID) error
	DeleteGroup(groupID groupservice.ID) error
	AddUserToGroup(groupId, userID groupservice.ID) error
	RemoveUserFromGroup(groupId, userID groupservice.ID) error

	// Reactions
	//
	AddReactionToDM(messageID reactionservice.ID, reaction messageservice.Reaction) error
	ChangeReactionToDM(messageID, userID reactionservice.ID, reaction messageservice.Reaction) error
	RemoveReactionFromDM(messageID, userID reactionservice.ID) error

	AddReactionToGroup(messageID, userID groupmessageservice.ID, reaction groupmessageservice.Reaction) error
	ChangeReactionToGroup(messageID groupmessageservice.ID, reaction groupmessageservice.Reaction) error
	RemoveReactionFromGroup(messageID, userID groupmessageservice.ID) error
}
