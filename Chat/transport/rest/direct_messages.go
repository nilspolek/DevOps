package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
)

func (rest *Rest) getDirectMessages(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	messages, err := (*(*rest).dms).GetMessages(messageservice.ID(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (rest *Rest) sendDirectMessage(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var message messageservice.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message.SenderID = messageservice.ID(id)
	err = (*(*rest).dms).SendMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) replaceDirectMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var message messageservice.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message.SenderID = messageservice.ID(id)
	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).dms).ReplaceMessage(messageservice.ID(messageId), message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (rest *Rest) deleteDirectMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	messageId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = (*(*rest).dms).DeleteMessage(messageservice.ID(messageId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
