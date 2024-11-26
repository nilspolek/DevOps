package dm_impl

import (
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	"github.com/nilspolek/DevOps/Chat/repo"
)

type svc struct {
	Repo repo.Repo
}

func New(repo repo.Repo) messageservice.DirectMessageService {
	return svc{Repo: repo}
}

func (s svc) GetMessages(userID messageservice.ID) ([]messageservice.Message, error) {
	return s.Repo.GetDirectMessages(userID)
}

func (s svc) SendMessage(msg messageservice.Message) error {
	return s.Repo.SendDirectMessage(msg)
}

func (s svc) ReplaceMessage(messageID messageservice.ID, msg messageservice.Message) error {
	return s.Repo.ReplaceDirecMessage(messageID, msg)
}

func (s svc) DeleteMessage(messageID messageservice.ID) error {
	return s.Repo.DeleteDirectMessage(messageID)
}
