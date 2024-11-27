package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
)

// Swagger API Documentation
// @Summary get all messages from a group
// @Description get all messages from a group.
// @Tags Group Messages
// @Accept  json
// @Produce  json
// @Param groupId path string true "groupId"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /group/{groupId}/messages [delete]
func (rest *Rest) getGroupMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	gId, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	gMessages, err := (*(*rest).gms).GetMessages(gId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gMessages)
}

// Swagger API Documentation
// @Summary send a message to a group
// @Description send a message to a group.
// @Tags Group Messages
// @Accept  json
// @Produce  json
// @Param groupId path string true "groupId"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /group/{groupId}/messages [post]
func (rest *Rest) sendGroupMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
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
	err = (*(*rest).gms).SendMessage(gId, gMessage, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
// @Summary replace a message in a group
// @Description replace a message in a group.
// @Tags Group Messages
// @Accept  json
// @Produce  json
// @Param groupId path string true "groupId"
// @Param messageId path string true "messageId"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /group/{groupId}/messages/{messageId} [put]
func (rest *Rest) replaceGroupMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
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
	err = (*(*rest).gms).ReplaceMessage(gId, gMessage, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
// @Summary replace a message in a group
// @Description replace a message in a group.
// @Tags Group Messages
// @Accept  json
// @Produce  json
// @Param groupId path string true "groupId"
// @Param messageId path string true "messageId"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /group/{groupId}/messages/{messageId} [delete]
func (rest *Rest) deleteGroupMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
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
	err = (*(*rest).gms).DeleteMessage(mId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
