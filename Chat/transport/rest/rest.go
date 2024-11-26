package rest

import (
	"github.com/gorilla/mux"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
	jwtservice "github.com/nilspolek/DevOps/Chat/jwt_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
)

type Rest struct {
	Router *mux.Router
	dms    *messageservice.DirectMessageService
	gms    *groupmessageservice.GroupMessageService
	gs     *groupservice.GroupService
	jwt    *jwtservice.JWTService
	rs     *reactionservice.ReactionService
}

func New(router *mux.Router, dms *messageservice.DirectMessageService, gms *groupmessageservice.GroupMessageService, gs *groupservice.GroupService, jwt *jwtservice.JWTService) *Rest {
	rest := &Rest{
		Router: router,
		dms:    dms,
		gms:    gms,
		gs:     gs,
		jwt:    jwt,
	}
	rest.RegisterRoutes()
	return rest
}

func (r *Rest) RegisterRoutes() {
	// Direct Messages
	//
	r.Router.HandleFunc("/messages", r.getDirectMessages).Methods("GET")
	r.Router.HandleFunc("/messages", r.sendDirectMessage).Methods("POST")
	r.Router.HandleFunc("/messages/{messageId}", r.replaceDirectMessage).Methods("PUT")
	r.Router.HandleFunc("/messages/{messageId}", r.deleteDirectMessage).Methods("DELETE")

	// Direct Message Reaction
	//
	r.Router.HandleFunc("/messages/{messageId}/react", r.addDirectMessageReaction).Methods("POST")
	r.Router.HandleFunc("/messages/{messageId}/react", r.changeDirectMessageReaction).Methods("PUT")
	r.Router.HandleFunc("/messages/{messageId}/react", r.deleteDirectMessageReaction).Methods("DELETE")

	// Group Messages
	//
	r.Router.HandleFunc("/group/messages", r.getGroupMessages).Methods("GET")
	r.Router.HandleFunc("/group/messages", r.sendGroupMessage).Methods("POST")
	r.Router.HandleFunc("/group/messages/{messageId}", r.replaceGroupMessage).Methods("PUT")
	r.Router.HandleFunc("/group/messages/{messageId}", r.deleteGroupMessage).Methods("DELETE")

	// Group Message Reaction
	//
	r.Router.HandleFunc("/group/message/{messageId}/reaction", r.addGroupMessageReaction).Methods("POST")
	r.Router.HandleFunc("/group/message/{messageId}/reaction", r.changeGroupReaction).Methods("PUT")
	r.Router.HandleFunc("/group/message/{messageId}/reaction", r.deleteGroupReaction).Methods("DELETE")

	// Groups
	//
	r.Router.HandleFunc("/groups", r.getGroups).Methods("GET")
	r.Router.HandleFunc("/groups", r.createGroup).Methods("POST")
	r.Router.HandleFunc("/groups/{groupId}", r.editGroup).Methods("PUT")
	r.Router.HandleFunc("/groups/{groupId}", r.deleteGroup).Methods("DELETE")
	r.Router.HandleFunc("/groups/{groupId}/users", r.addUserToGroup).Methods("POST")
	r.Router.HandleFunc("/groups/{groupId}/users/{userId}", r.removeUserFromGroup).Methods("DELETE")
}
