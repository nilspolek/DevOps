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
	msgId      uuid.UUID
)

func TestNew(t *testing.T) {
	repositories, err := mongodb.New()
	if err != nil {
		t.Error(err)
	}
	dm = dm_impl.New(repositories)
}

func TestSendMessages(t *testing.T) {
	var (
		err error
	)
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
	msgId, err = dm.SendMessage(message, user)
	if err != nil {
		t.Error(err)
	}
}

func TestGetMessages(t *testing.T) {
	msgs, err := dm.GetMessages(receiverID, user)
	if err != nil {
		t.Error(err)
	}
	passed := false
	for _, msg := range msgs {
		if msg.Id == msgId {
			passed = true
		}
	}
	if !passed {
		t.Fatal("Message was not in messages")
	}
	if len(msgs) == 0 {
		t.Fatal("No messages found")
	}
}

func TestReplaceMessage(t *testing.T) {
	err := dm.ReplaceMessage(msgId, messageservice.Message{Content: "some other content"}, user)
	if err != nil {
		t.Fatal(err)
	}
	msgs, err := dm.GetMessages(user, user)
	for _, msg := range msgs {
		if msg.Id == msgId && msg.Content != "some other content" {
			t.Fatalf("Message was not replaced (%s)", msg.Content)
		}
	}
}

func TestDeleteMessage(t *testing.T) {
	err := dm.DeleteMessage(msgId, user)
	if err != nil {
		t.Fatal(err)
	}
	msgs, err := dm.GetMessages(receiverID, user)
	for _, msg := range msgs {
		if msg.Id == msgId {
			t.Fatal("Message was not deleted")
		}
	}
}
