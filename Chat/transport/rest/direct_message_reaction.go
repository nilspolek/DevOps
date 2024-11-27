package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
)

// Swagger API Documentation
// @Summary adds a reaction to a direct message
// @Description adds a reaction to a direct message.
// @Tags Direct Message Reaction
// @Accept  json
// @Produce  json
// @Param messageId path string true "The id of the message"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /messages/{messageId}/react [get]
func (rest *Rest) addDirectMessageReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var reaction messageservice.Reaction
	err = json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reaction.Sender = id
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).AddReactionToDM(messageId, reaction, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
// @Summary change a reaction to a direct message
// @Description changes a reaction to a direct message.
// @Tags Direct Message Reaction
// @Accept  json
// @Produce  json
// @Param messageId path string true "The id of the message"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /messages/{messageId}/react [put]
func (rest *Rest) changeDirectMessageReaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var reaction messageservice.Reaction
	err = json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reaction.Sender = id
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).rs).ChangeReactionToDM(messageId, reaction, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
// @Summary delete a reaction to a direct message
// @Description delete a reaction to a direct message.
// @Tags Direct Message Reaction
// @Accept  json
// @Produce  json
// @Param messageId path string true "The id of the message"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /messages/{messageId}/react [delete]
func (rest *Rest) deleteDirectMessageReaction(w http.ResponseWriter, r *http.Request) {
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
	err = (*(*rest).rs).RemoveReactionFromDM(messageId, id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
