package gmlog

import (
	"time"

	"github.com/google/uuid"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	"github.com/nilspolek/goLog"
)

type svc struct {
	next *groupmessageservice.GroupMessageService
}

func New(next *groupmessageservice.GroupMessageService) groupmessageservice.GroupMessageService {
	svc := svc{
		next: next,
	}
	return svc
}

func (s svc) GetMessages(groupID, authUser uuid.UUID) (msgs []groupmessageservice.Message, err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Get Group Message | took: %s | %v", time.Since(tm), msgs)
	}(time.Now())
	msgs, err = (*s.next).GetMessages(groupID, authUser)
	return
}

func (s svc) SendMessage(groupID uuid.UUID, msg groupmessageservice.Message, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Send Group Message | took: %s | %v", time.Since(tm), msg)
	}(time.Now())
	err = (*s.next).SendMessage(groupID, msg, authUser)
	return
}

func (s svc) ReplaceMessage(messageID uuid.UUID, msg groupmessageservice.Message, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Replace Group Message | took: %s | %v", time.Since(tm))
	}(time.Now())
	err = (*s.next).ReplaceMessage(messageID, msg, authUser)
	return
}

func (s svc) DeleteMessage(messageID, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Delete Group Message | took: %s | %v", time.Since(tm))
	}(time.Now())
	err = (*s.next).DeleteMessage(messageID, authUser)
	return
}
