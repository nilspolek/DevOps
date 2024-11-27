package rlog

import (
	"time"

	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
	"github.com/nilspolek/goLog"
)

type svc struct {
	next reactionservice.ReactionService
}

func New(next reactionservice.ReactionService, prefix string) reactionservice.ReactionService {
	svc := svc{
		next: next,
	}
	return &svc
}

func (s *svc) AddReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Add Reaction to Direct Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.AddReactionToDM(messageID, reaction, authUser)
	return nil
}

func (s *svc) ChangeReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Changed Reaction from Direct Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.ChangeReactionToDM(messageID, reaction, authUser)
	return nil
}

func (s *svc) RemoveReactionFromDM(messageID, userID, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Removed Reaction from Direct Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.RemoveReactionFromDM(messageID, userID, authUser)
	return nil
}

func (s *svc) AddReactionToGroup(messageID, userId uuid.UUID, reaction groupmessageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Add Reaction to Group Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.AddReactionToGroup(messageID, userId, reaction, authUser)
	return nil
}

func (s *svc) ChangeReactionToGroup(messageID uuid.UUID, reaction groupmessageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Change Reaction from Group Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.ChangeReactionToGroup(messageID, reaction, authUser)
	return nil
}

func (s *svc) RemoveReactionFromGroup(messageID, userID, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Removed Reaction from Group Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.RemoveReactionFromGroup(messageID, userID, authUser)
	return nil
}
