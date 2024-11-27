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
	r.Router.HandleFunc("/messages", r.jwtMiddleware(r.getDirectMessages)).Methods("GET")
	r.Router.HandleFunc("/messages", r.jwtMiddleware(r.sendDirectMessage)).Methods("POST")
	r.Router.HandleFunc("/messages/{messageId}", r.jwtMiddleware(r.replaceDirectMessage)).Methods("PUT")
	r.Router.HandleFunc("/messages/{messageId}", r.jwtMiddleware(r.deleteDirectMessage)).Methods("DELETE")

	// Direct Message Reaction
	//
	r.Router.HandleFunc("/messages/{messageId}/react", r.jwtMiddleware(r.addDirectMessageReaction)).Methods("POST")
	r.Router.HandleFunc("/messages/{messageId}/react", r.jwtMiddleware(r.changeDirectMessageReaction)).Methods("PUT")
	r.Router.HandleFunc("/messages/{messageId}/react", r.jwtMiddleware(r.deleteDirectMessageReaction)).Methods("DELETE")

	// Group Messages
	//
	r.Router.HandleFunc("/group/messages", r.jwtMiddleware(r.getGroupMessages)).Methods("GET")
	r.Router.HandleFunc("/group/messages", r.jwtMiddleware(r.sendGroupMessage)).Methods("POST")
	r.Router.HandleFunc("/group/messages/{messageId}", r.jwtMiddleware(r.replaceGroupMessage)).Methods("PUT")
	r.Router.HandleFunc("/group/messages/{messageId}", r.jwtMiddleware(r.deleteGroupMessage)).Methods("DELETE")

	// Group Message Reaction
	//
	r.Router.HandleFunc("/group/message/{messageId}/reaction", r.jwtMiddleware(r.addGroupMessageReaction)).Methods("POST")
	r.Router.HandleFunc("/group/message/{messageId}/reaction", r.jwtMiddleware(r.changeGroupReaction)).Methods("PUT")
	r.Router.HandleFunc("/group/message/{messageId}/reaction", r.jwtMiddleware(r.deleteGroupReaction)).Methods("DELETE")

	// Groups
	//
	r.Router.HandleFunc("/groups", r.jwtMiddleware(r.getGroups)).Methods("GET")
	r.Router.HandleFunc("/groups", r.jwtMiddleware(r.createGroup)).Methods("POST")
	r.Router.HandleFunc("/groups/{groupId}", r.jwtMiddleware(r.editGroup)).Methods("PUT")
	r.Router.HandleFunc("/groups/{groupId}", r.jwtMiddleware(r.deleteGroup)).Methods("DELETE")
	r.Router.HandleFunc("/groups/{groupId}/users", r.jwtMiddleware(r.addUserToGroup)).Methods("POST")
	r.Router.HandleFunc("/groups/{groupId}/users/{userId}", r.jwtMiddleware(r.removeUserFromGroup)).Methods("DELETE")
}
