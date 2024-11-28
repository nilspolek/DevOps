package dm_impl

import (
	"time"

	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	"github.com/nilspolek/DevOps/Chat/repo"
)

type svc struct {
	Repo repo.Repo
}

func New(repo repo.Repo) messageservice.DirectMessageService {
	return svc{Repo: repo}
}

func (s svc) GetMessages(userID, authUser uuid.UUID) ([]messageservice.Message, error) {
	return s.Repo.GetDirectMessages(userID)
}

func (s svc) SendMessage(msg messageservice.Message, authUser uuid.UUID) (uuid.UUID, error) {
	msg.Id = uuid.New()
	msg.SenderID = authUser
	msg.Timestamp = time.Now()
	return msg.Id, s.Repo.SendDirectMessage(msg)
}

func (s svc) ReplaceMessage(messageID uuid.UUID, msg messageservice.Message, authUser uuid.UUID) error {
	msg.SenderID = authUser
	return s.Repo.ReplaceDirecMessage(messageID, msg)
}

func (s svc) DeleteMessage(messageID, authUser uuid.UUID) error {
	return s.Repo.DeleteDirectMessage(messageID)
}
