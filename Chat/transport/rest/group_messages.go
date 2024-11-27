package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
)

func (rest *Rest) getGroupMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	gId, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	gMessages, err := (*(*rest).gms).GetMessages(gId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gMessages)
}

func (rest *Rest) sendGroupMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	gId, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var gMessage groupmessageservice.Message
	err = json.NewDecoder(r.Body).Decode(&gMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gms).SendMessage(gId, gMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) replaceGroupMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	gId, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var gMessage groupmessageservice.Message
	err = json.NewDecoder(r.Body).Decode(&gMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gms).ReplaceMessage(gId, gMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) deleteGroupMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	_, err = uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	mId, err := uuid.Parse(vars["messageId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gms).DeleteMessage(mId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
