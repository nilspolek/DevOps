package rprometheus

import (
	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
	"github.com/prometheus/client_golang/prometheus"
)

type svc struct {
	next reactionservice.ReactionService

	addReactionDMCounter    prometheus.Counter
	changeReactionDMCounter prometheus.Counter
	removeReactionDMCounter prometheus.Counter

	addReactionGroupCounter    prometheus.Counter
	changeReactionGroupCounter prometheus.Counter
	removeReactionGroupCounter prometheus.Counter
	errorReactionCounter       prometheus.Counter
}

func New(next reactionservice.ReactionService, prefix string) (reactionservice.ReactionService, error) {
	svc := svc{
		addReactionDMCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prefix + "add_reaction_dm_total",
			Help: "Total number of reactions added to direct messages",
		}),
		changeReactionDMCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prefix + "change_reaction_dm_total",
			Help: "Total number of reactions changed in direct messages",
		}),
		removeReactionDMCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prefix + "remove_reaction_dm_total",
			Help: "Total number of reactions removed from direct messages",
		}),
		addReactionGroupCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prefix + "add_reaction_group_total",
			Help: "Total number of reactions added to group messages",
		}),
		changeReactionGroupCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prefix + "change_reaction_group_total",
			Help: "Total number of reactions changed in group messages",
		}),
		removeReactionGroupCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prefix + "remove_reaction_group_total",
			Help: "Total number of reactions removed from group messages",
		}),
		errorReactionCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prefix + "error_reaction_total",
			Help: "Total number of errors in the reaction service",
		}),
	}
	err := prometheus.Register(svc.addReactionDMCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.changeReactionDMCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.removeReactionDMCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.addReactionGroupCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.changeReactionGroupCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.removeReactionGroupCounter)
	return &svc, err
}

func (s *svc) AddReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorReactionCounter.Inc()
		}
	}()
	err = s.next.AddReactionToDM(messageID, reaction, authUser)
	s.addReactionDMCounter.Inc()
	return nil
}

func (s *svc) ChangeReactionToDM(messageID uuid.UUID, reaction messageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorReactionCounter.Inc()
		}
	}()
	err = s.next.ChangeReactionToDM(messageID, reaction, authUser)
	s.changeReactionDMCounter.Inc()
	return nil
}

func (s *svc) RemoveReactionFromDM(messageID, userID, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorReactionCounter.Inc()
		}
	}()
	err = s.next.RemoveReactionFromDM(messageID, userID, authUser)
	s.removeReactionDMCounter.Inc()
	return nil
}

func (s *svc) AddReactionToGroup(messageID, userId uuid.UUID, reaction groupmessageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorReactionCounter.Inc()
		}
	}()
	err = s.next.AddReactionToGroup(messageID, userId, reaction, authUser)
	s.addReactionGroupCounter.Inc()
	return nil
}

func (s *svc) ChangeReactionToGroup(messageID uuid.UUID, reaction groupmessageservice.Reaction, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorReactionCounter.Inc()
		}
	}()
	err = s.next.ChangeReactionToGroup(messageID, reaction, authUser)
	s.changeReactionGroupCounter.Inc()
	return nil
}

func (s *svc) RemoveReactionFromGroup(messageID, userID, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorReactionCounter.Inc()
		}
	}()
	err = s.next.RemoveReactionFromGroup(messageID, userID, authUser)
	s.removeReactionGroupCounter.Inc()
	return nil
}
