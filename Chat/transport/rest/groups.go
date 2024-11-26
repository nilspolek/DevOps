package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
)

func (rest *Rest) getGroups(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	id, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	groups, err := (*(*rest).gs).GetAllGroups(groupservice.ID(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func (rest *Rest) createGroup(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var group groupservice.Group
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupId, err := (*(*rest).gs).CreateGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupId)
}

func (rest *Rest) editGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var group groupservice.Group
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gID, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gs).EditGroup(group, groupservice.ID(gID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) deleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	gID, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gs).DeleteGroup(groupservice.ID(gID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) addUserToGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	gID, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	uID, err := uuid.Parse(vars["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gs).AddUserToGroup(groupservice.ID(gID), groupservice.ID(uID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rest *Rest) removeUserFromGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authHeader := r.Header.Get("Authorization")
	_, err := (*(*rest).jwt).ValiadteToken(authHeader)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	gID, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	uID, err := uuid.Parse(vars["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gs).RemoveUserFromGroup(groupservice.ID(gID), groupservice.ID(uID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
