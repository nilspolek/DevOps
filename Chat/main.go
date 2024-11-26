package main

import (
	"log"

	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
	gimpl "github.com/nilspolek/DevOps/Chat/group_service/g_impl"
	"github.com/nilspolek/DevOps/Chat/repo/mongodb"
)

func main() {
	repo, err := mongodb.New()
	if err != nil {
		log.Fatal(err)
	}
	g_impl := gimpl.New(repo)
	g_impl.CreateGroup(groupservice.Group{})
}
