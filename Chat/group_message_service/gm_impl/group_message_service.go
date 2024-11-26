package gmimpl

import (
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	"github.com/nilspolek/DevOps/Chat/repo"
)

type svc struct {
	Repo repo.Repo
}

func New(repo repo.Repo) groupmessageservice.GroupMessageService {
	return svc{Repo: repo}
}

func (s svc) GetMessages(groupID groupmessageservice.ID) ([]groupmessageservice.Message, error) {
	return s.Repo.GetGroupMessages(groupID)
}

func (s svc) SendMessage(gid groupmessageservice.ID, msg groupmessageservice.Message) error {
	return s.Repo.SendMessageToGroup(msg.GroupId, msg)
}

func (s svc) ReplaceMessage(messageID groupmessageservice.ID, msg groupmessageservice.Message) error {
	return s.Repo.ReplaceGroupMessage(messageID, msg)
}

func (s svc) DeleteMessage(messageID groupmessageservice.ID) error {
	return s.Repo.DeleteGroupMessage(messageID)
}
