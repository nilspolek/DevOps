package rimpl

import (
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

func (s *svc) AddReactionToDM(messageID reactionservice.ID, reaction messageservice.Reaction) error {
	return s.repo.AddReactionToDM(messageID, reaction)
}

func (s *svc) ChangeReactionToDM(messageID reactionservice.ID, reaction messageservice.Reaction) error {
	return s.repo.ChangeReactionToDM(messageID, reactionservice.ID(reaction.Sender), reaction)
}

func (s *svc) RemoveReactionFromDM(messageID, userId reactionservice.ID) error {
	return s.repo.RemoveReactionFromDM(messageID, userId)
}

func (s *svc) AddReactionToGroup(messageID, userId reactionservice.ID, reaction groupmessageservice.Reaction) error {
	return s.repo.AddReactionToGroup(groupmessageservice.ID(messageID), groupmessageservice.ID(userId), reaction)
}

func (s *svc) ChangeReactionToGroup(messageID reactionservice.ID, reaction groupmessageservice.Reaction) error {
	return s.repo.ChangeReactionToGroup(groupmessageservice.ID(messageID), reaction)
}

func (s *svc) RemoveReactionFromGroup(messageID, userID reactionservice.ID) error {
	return s.repo.RemoveReactionFromGroup(groupmessageservice.ID(messageID), groupmessageservice.ID(userID))
}
