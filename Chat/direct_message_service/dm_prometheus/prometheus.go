package dmprometheus

import (
	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	"github.com/prometheus/client_golang/prometheus"
)

type svc struct {
	next                  messageservice.DirectMessageService
	getMessageCounter     prometheus.Counter
	sendMessageCounter    prometheus.Counter
	replaceMessageCounter prometheus.Counter
	deleteMessageCounter  prometheus.Counter
	errorCounter          prometheus.Counter
}

func New(next messageservice.DirectMessageService, prefix string) (messageservice.DirectMessageService, error) {
	svc := svc{
		next:                  next,
		getMessageCounter:     prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_get_direct_message_counter", Help: "Number of get messages calls"}),
		sendMessageCounter:    prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_send_direct_message_counter", Help: "Number of send messages calls"}),
		replaceMessageCounter: prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_replace_direct_message_counter", Help: "Number of replace messages calls"}),
		deleteMessageCounter:  prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_delete_direct_message_counter", Help: "Number of delete messages calls"}),
		errorCounter:          prometheus.NewCounter(prometheus.CounterOpts{Name: prefix + "_error_direct_message_counter", Help: "Number of errors"}),
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

func (s svc) GetMessages(userID uuid.UUID) (msgs []messageservice.Message, err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.getMessageCounter.Inc()
	msgs, err = s.next.GetMessages(userID)
	return
}

func (s svc) SendMessage(msg messageservice.Message) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.sendMessageCounter.Inc()
	err = s.next.SendMessage(msg)
	return
}

func (s svc) ReplaceMessage(messageID uuid.UUID, msg messageservice.Message) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.replaceMessageCounter.Inc()
	err = s.next.ReplaceMessage(messageID, msg)
	return
}

func (s svc) DeleteMessage(messageID uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			s.errorCounter.Inc()
		}
	}()
	s.deleteMessageCounter.Inc()
	err = s.next.DeleteMessage(messageID)
	return
}
