package dmlog

import (
	"time"

	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	"github.com/nilspolek/goLog"
)

type svc struct {
	next messageservice.DirectMessageService
}

func New(next messageservice.DirectMessageService) messageservice.DirectMessageService {
	return svc{next: next}
}

func (s svc) GetMessages(userID, authUser uuid.UUID) (msgs []messageservice.Message, err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Get Direct Message | took: %s | %v", time.Since(tm), msgs)
	}(time.Now())
	msgs, err = s.next.GetMessages(userID, authUser)
	return
}

func (s svc) SendMessage(msg messageservice.Message, authUser uuid.UUID) (id uuid.UUID, err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Send Direct Message | took: %s | %v", time.Since(tm), msg)
	}(time.Now())
	id, err = s.next.SendMessage(msg, authUser)
	return
}

func (s svc) ReplaceMessage(messageID uuid.UUID, msg messageservice.Message, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Replace Direct Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.ReplaceMessage(messageID, msg, authUser)
	return
}

func (s svc) DeleteMessage(messageID, authUser uuid.UUID) (err error) {
	defer func(tm time.Time) {
		if err != nil {
			goLog.Error("%v", err)
		}
		goLog.Info("Delete Direct Message | took: %s", time.Since(tm))
	}(time.Now())
	err = s.next.DeleteMessage(messageID, authUser)
	return
}
