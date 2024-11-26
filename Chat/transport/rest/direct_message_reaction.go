package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
)

func (rest *Rest) addDirectMessageReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var reaction messageservice.Reaction
	err = json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reaction.Sender = messageservice.ID(id)
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).AddReactionToDM(reactionservice.ID(messageId), reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) changeDirectMessageReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var reaction messageservice.Reaction
	err = json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reaction.Sender = messageservice.ID(id)
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).ChangeReactionToDM(reactionservice.ID(messageId), reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) deleteDirectMessageReaction(w http.ResponseWriter, r *http.Request) {
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
	err = (*(*rest).rs).RemoveReactionFromDM(reactionservice.ID(messageId), reactionservice.ID(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
