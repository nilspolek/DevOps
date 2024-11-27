package glog

import (
	"time"

	"github.com/google/uuid"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
	"github.com/nilspolek/goLog"
)

type svc struct {
	next *groupservice.GroupService
}

func New(next *groupservice.GroupService) groupservice.GroupService {
	svc := svc{
		next: next,
	}
	return svc
}

func (s svc) GetAllGroups(userId, authUser uuid.UUID) (gps []groupservice.Group, err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Get Group | took: %s | %v", time.Since(tm), gps)
	}(time.Now())
	gps, err = (*s.next).GetAllGroups(userId, authUser)
	return
}

func (s svc) CreateGroup(group groupservice.Group, authUser uuid.UUID) (id uuid.UUID, err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Create Group | took: %s | %v", time.Since(tm), id)
	}(time.Now())
	id, err = (*s.next).CreateGroup(group, authUser)
	return
}

func (s svc) EditGroup(group groupservice.Group, id, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Edit Group | took: %s", time.Since(tm))
	}(time.Now())
	err = (*s.next).EditGroup(group, id, authUser)
	return
}

func (s svc) DeleteGroup(id, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Delete Group | took: %s ", time.Since(tm))
	}(time.Now())
	err = (*s.next).DeleteGroup(id, authUser)
	return
}

func (s svc) AddUserToGroup(groupId, userId, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Add User to Group | took: %s", time.Since(tm))
	}(time.Now())
	err = (*s.next).AddUserToGroup(groupId, userId, authUser)
	return
}

func (s svc) RemoveUserFromGroup(groupId, userId, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Remove User Group | took: %s", time.Since(tm))
	}(time.Now())
	err = (*s.next).RemoveUserFromGroup(groupId, userId, authUser)
	return
}
