package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	"github.com/nilspolek/DevOps/Chat/direct_message_service/dm_impl"
	"github.com/nilspolek/DevOps/Chat/repo/mongodb"
)

func main() {
	repo, err := mongodb.New()
	if err != nil {
		log.Fatal(err)
	}

	dms := dm_impl.New(repo)
	id := messageservice.ID(uuid.MustParse("25b154f9-5a8b-49c1-b676-f4a821838fc1"))
	msgs, err := dms.GetMessages(id)

	fmt.Printf("%+v", msgs)
}
