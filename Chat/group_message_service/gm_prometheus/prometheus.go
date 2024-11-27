package gmprometheus

import (
	"github.com/google/uuid"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	"github.com/prometheus/client_golang/prometheus"
)

type svc struct {
	next                  groupmessageservice.GroupMessageService
	getMessageCounter     prometheus.Counter
	sendMessageCounter    prometheus.Counter
	replaceMessageCounter prometheus.Counter
	deleteMessageCounter  prometheus.Counter
	errorCounter          prometheus.Counter
}

func New(next groupmessageservice.GroupMessageService, prefix string) (groupmessageservice.GroupMessageService, error) {
	svc := svc{
		next:                  next,
		getMessageCounter:     prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_get_group_message_counter", Help: "Number of get messages calls"}),
		sendMessageCounter:    prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_send_group_message_counter", Help: "Number of send messages calls"}),
		replaceMessageCounter: prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_replace_group_message_counter", Help: "Number of replace messages calls"}),
		deleteMessageCounter:  prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_delete-group_message_counter", Help: "Number of delete messages calls"}),
		errorCounter:          prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_error-group_counter", Help: "Number of errors"}),
	}
	err := prometheus.Register(svc.getMessageCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.sendMessageCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.replaceMessageCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.deleteMessageCounter)
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(svc.errorCounter)
	return svc, err
}

func (s svc) GetMessages(groupID, authUser uuid.UUID) (msgs []groupmessageservice.Message, err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.getMessageCounter.Inc()
	msgs, err = s.next.GetMessages(groupID, authUser)
	return
}

func (s svc) SendMessage(groupID uuid.UUID, msg groupmessageservice.Message, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.sendMessageCounter.Inc()
	err = s.next.SendMessage(groupID, msg, authUser)
	return
}

func (s svc) ReplaceMessage(messageID uuid.UUID, msg groupmessageservice.Message, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.replaceMessageCounter.Inc()
	err = s.next.ReplaceMessage(messageID, msg, authUser)
	return
}

func (s svc) DeleteMessage(messageID, authUser uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.deleteMessageCounter.Inc()
	err = s.next.DeleteMessage(messageID, authUser)
	return
}
