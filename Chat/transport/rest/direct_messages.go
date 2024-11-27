package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
)

// Swagger API Documentation
// @Summary get all direct messages a user has access to
// @Description get all direct messages a user has access to.
// @Tags Direct Messages
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /messages [get]
func (rest *Rest) getDirectMessages(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	messages, err := (*(*rest).dms).GetMessages(id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// Swagger API Documentation
// @Summary send a direct message
// @Description send a direct message to a user.
// @Tags Direct Messages
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /messages [post]
func (rest *Rest) sendDirectMessage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var message messageservice.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message.SenderID = id
	err = (*(*rest).dms).SendMessage(message, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
// @Summary replace a direct message
// @Description replace a direct message.
// @Tags Direct Messages
// @Accept  json
// @Produce  json
// @Param messageId path string true "messageId"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /messages/{messageId} [put]
func (rest *Rest) replaceDirectMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var message messageservice.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message.SenderID = id
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).dms).ReplaceMessage(messageId, message, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
// @Summary delete a direct message
// @Description delete a direct message.
// @Tags Direct Messages
// @Accept  json
// @Produce  json
// @Param messageId path string true "messageId"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /messages/{messageId} [delete]
func (rest *Rest) deleteDirectMessage(w http.ResponseWriter, r *http.Request) {
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
	err = (*(*rest).dms).DeleteMessage(messageId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
