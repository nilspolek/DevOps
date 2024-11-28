package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/nilspolek/DevOps/Chat/direct_message_service/dm_impl"
	"github.com/nilspolek/DevOps/Chat/repo/mongodb"
	"testing"
)

func TestName(t *testing.T) {
	repository, err := mongodb.New()
	if err != nil {
		panic(err)
	}
	dms := dm_impl.New(repository)
	msgs, err := dms.GetMessages(uuid.New(), uuid.New())
	fmt.Println(msgs)
	t.Error()

}
