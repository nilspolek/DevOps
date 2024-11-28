package gprometheus

import (
	"github.com/google/uuid"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
	"github.com/prometheus/client_golang/prometheus"
)

type svc struct {
	next                groupservice.GroupService
	getAllGroupsCounter prometheus.Counter
	createGroupCounter  prometheus.Counter
	editGroupCounter    prometheus.Counter
	deleteGroupCounter  prometheus.Counter
	addUserCounter      prometheus.Counter
	removeUserCounter   prometheus.Counter
	errorCounter        prometheus.Counter
}

func New(next groupservice.GroupService, prefic string) (groupservice.GroupService, error) {
	svc := svc{
		next:                next,
		getAllGroupsCounter: prometheus.NewCounter(prometheus.CounterOpts{Name: prefic + "_get_all_groups_counter", Help: "Number of get all groups calls"}),
		createGroupCounter:  prometheus.NewCounter(prometheus.CounterOpts{Name: prefic + "_create_group_counter", Help: "Number of create group calls"}),
		editGroupCounter:    prometheus.NewCounter(prometheus.CounterOpts{Name: prefic + "_edit_group_counter", Help: "Number of edit group calls"}),
		deleteGroupCounter:  prometheus.NewCounter(prometheus.CounterOpts{Name: prefic + "_delete_group_counter", Help: "Number of delete group calls"}),
		addUserCounter:      prometheus.NewCounter(prometheus.CounterOpts{Name: prefic + "_add_user_counter", Help: "Number of add user calls"}),
		removeUserCounter:   prometheus.NewCounter(prometheus.CounterOpts{Name: prefic + "_remove_user_counter", Help: "Number of remove user calls"}),
		errorCounter:        prometheus.NewCounter(prometheus.CounterOpts{Name: prefic + "_error_counter", Help: "Number of errors"}),
	}
	err := prometheus.Register(svc.getAllGroupsCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.createGroupCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.editGroupCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.deleteGroupCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.addUserCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.removeUserCounter)
	if err != nil {
		return nil, err
	}
	return &svc, prometheus.Register(svc.errorCounter)
}

func (s svc) GetAllGroups(userId, authUser uuid.UUID) (gps []groupservice.Group, err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.getAllGroupsCounter.Inc()
	gps, err = s.next.GetAllGroups(userId, authUser)
	return
}

func (s svc) CreateGroup(group groupservice.Group, authUser uuid.UUID) (id uuid.UUID, err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.createGroupCounter.Inc()
	id, err = s.next.CreateGroup(group, authUser)
	return
}

func (s svc) EditGroup(group groupservice.Group, id, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.editGroupCounter.Inc()
	err = s.next.EditGroup(group, id, authUser)
	return
}

func (s svc) DeleteGroup(id, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.deleteGroupCounter.Inc()
	err = s.next.DeleteGroup(id, authUser)
	return
}

func (s svc) AddUserToGroup(groupId, userId, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.addUserCounter.Inc()
	err = s.next.AddUserToGroup(groupId, userId, authUser)
	return
}

func (s svc) RemoveUserFromGroup(groupId, userId, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.removeUserCounter.Inc()
	err = s.next.RemoveUserFromGroup(groupId, userId, authUser)
	return
}
