package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
)

// Swagger API Documentation
//
//	@Summary		add a reaction to a message from a group
//	@Description	add a reaction to a message from a group.
//	@Tags			Group Message Reactions
//	@Accept			json
//	@Produce		json
//
//	@Param			reaction	body		groupmessageservice.Reaction	true	"reaction"
//
//	@Param			messageId	path		string							true	"messageId"
//	@Success		200			{object}	string
//	@Failure		400			{object}	string
//	@Router			/group/message/{messageId}/reaction [post]
func (rest *Rest) addGroupMessageReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
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
	reaction.Sender.Id = id
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).AddReactionToGroup(messageId, id, reaction, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
//
//	@Summary		change a reaction to a message from a group
//	@Description	change a reaction to a message from a group.
//	@Tags			Group Message Reactions
//	@Accept			json
//	@Produce		json
//	@Param			messageId	path		string							true	"messageId"
//
//	@Param			reaction	body		groupmessageservice.Reaction	true	"reaction"
//
//	@Success		200			{object}	string
//	@Failure		400			{object}	string
//	@Router			/group/message/{messageId}/reaction [put]
func (rest *Rest) changeGroupReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
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

	reaction.Sender.Id = id
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).ChangeReactionToGroup(messageId, reaction, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
//
//	@Summary		delete a reaction to a message from a group
//	@Description	delete a reaction to a message from a group.
//	@Tags			Group Message Reactions
//	@Accept			json
//	@Produce		json
//	@Param			messageId	path		string	true	"messageId"
//	@Success		200			{object}	string
//	@Failure		400			{object}	string
//	@Router			/group/message/{messageId}/reaction [delete]
func (rest *Rest) deleteGroupReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).RemoveReactionFromGroup(messageId, id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
