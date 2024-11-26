package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
)

func (rest *Rest) addGroupMessageReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var reaction groupmessageservice.Reaction
	err = json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reaction.Sender.Id = groupmessageservice.ID(id)
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).AddReactionToGroup(reactionservice.ID(messageId), reactionservice.ID(id), reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) changeGroupReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var reaction groupmessageservice.Reaction
	err = json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reaction.Sender.Id = groupmessageservice.ID(id)
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).ChangeReactionToGroup(reactionservice.ID(messageId), reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) deleteGroupReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).RemoveReactionFromGroup(reactionservice.ID(messageId), reactionservice.ID(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
