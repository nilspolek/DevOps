package gmtest

import (
	"testing"

	"github.com/google/uuid"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	gmimpl "github.com/nilspolek/DevOps/Chat/group_message_service/gm_impl"
	"github.com/nilspolek/DevOps/Chat/repo/mongodb"
)

var (
	gm      groupmessageservice.GroupMessageService
	userId  = uuid.New()
	groupId = uuid.New()
	msgId   uuid.UUID
)

func TestNew(t *testing.T) {
	repository, err := mongodb.New()
	if err != nil {
		t.Error(err)
	}
	gm = gmimpl.New(repository)
}

func TestSendMessages(t *testing.T) {
	var err error
	msg := groupmessageservice.Message{
		Content: "Some content",
		GroupId: groupId,
	}
	msgId, err = gm.SendMessage(groupId, msg, userId)
	if err != nil {
		t.Error(err)
	}
}

func TestGetMessage(t *testing.T) {
	msgs, err := gm.GetMessages(groupId, userId)
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
		t.Errorf("Message not found %v", msgs)
	}
}

func TestReplaceMessage(t *testing.T) {
	err := gm.ReplaceMessage(msgId, groupmessageservice.Message{
		Content: "Some other content",
	}, userId)

	if err != nil {
		t.Error(err)
	}
	msgs, err := gm.GetMessages(groupId, userId)
	if err != nil {
		t.Error(err)
	}
	for _, msg := range msgs {
		if msg.Id == msgId && msg.Content != "Some other content" {
			t.Fatalf("Message was not replaced (%s)", msg.Content)
		}
	}
}

func TestDeleteMessage(t *testing.T) {
	err := gm.DeleteMessage(msgId, userId)
	if err != nil {
		t.Error(err)
	}
	msgs, err := gm.GetMessages(groupId, userId)
	if err != nil {
		t.Error(err)
	}
	for _, msg := range msgs {
		if msg.Id == msgId {
			t.Fatalf("Message was not deleted (%s)", msg.Content)
		}
	}
}
