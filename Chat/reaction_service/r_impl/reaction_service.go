package rimpl

import (
	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
	"github.com/nilspolek/DevOps/Chat/repo"
)

type svc struct {
	repo repo.Repo
}

func New(repo repo.Repo) reactionservice.ReactionService {
	return &svc{repo: repo}
}

func (s *svc) AddReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction) error {
	return s.repo.AddReactionToDM(messageID, reaction)
}

func (s *svc) ChangeReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction) error {
	return s.repo.ChangeReactionToDM(messageID, reaction.Sender, reaction)
}

func (s *svc) RemoveReactionFromDM(messageID, userId uuid.UUID) error {
	return s.repo.RemoveReactionFromDM(messageID, userId)
}

func (s *svc) AddReactionToGroup(messageID, userId uuid.UUID, reaction groupmessageservice.Reaction) error {
	return s.repo.AddReactionToGroup(messageID, userId, reaction)
}

func (s *svc) ChangeReactionToGroup(messageID uuid.UUID, reaction groupmessageservice.Reaction) error {
	return s.repo.ChangeReactionToGroup(messageID, reaction)
}

func (s *svc) RemoveReactionFromGroup(messageID, userID uuid.UUID) error {
	return s.repo.RemoveReactionFromGroup(messageID, userID)
}
