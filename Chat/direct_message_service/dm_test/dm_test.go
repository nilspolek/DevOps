package dmtestgo

import (
	"testing"

	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	"github.com/nilspolek/DevOps/Chat/direct_message_service/dm_impl"
	"github.com/nilspolek/DevOps/Chat/repo/mongodb"
)

var (
	receiverID = uuid.New()
	user       = uuid.New()
	dm         messageservice.DirectMessageService
)

func TestNew(t *testing.T) {
	repositories, err := mongodb.New()
	if err != nil {
		t.Error(err)
	}
	dm = dm_impl.New(repositories)
}

func TestSendMessages(t *testing.T) {
	message := messageservice.Message{
		Content:    "Some test Content",
		ReceiverID: receiverID,
		ImageUrl:   "https://www.google.com",
		Reactions: []messageservice.Reaction{
			{
				Sender:   user,
				Reaction: "like",
			}, {
				Sender:   receiverID,
				Reaction: "dislike",
			},
		},
	}
	dm.SendMessage(message, user)
}

func TestGetMessages(t *testing.T) {
	msgs, err := dm.GetMessages(receiverID, user)
	if err != nil {
		t.Error(err)
	}
	t.Log(msgs)
}
