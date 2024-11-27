package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
)

// Swagger API Documentation
//
//	@Summary		get all groups a user can access
//	@Description	get all groups a user can access.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Router			/groups [get]
func (rest *Rest) getGroups(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	groups, err := (*(*rest).gs).GetAllGroups(id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

// Swagger API Documentation
//
//	@Summary		create a group
//	@Description	create a group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//
//	@Param			group	body		groupservice.Group	true	"Group"
//
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Router			/groups [post]
func (rest *Rest) createGroup(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.Header.Get("userId"))
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

	groupId, err := (*(*rest).gs).CreateGroup(group, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupId)
}

// Swagger API Documentation
//
//	@Summary		edit a Group
//	@Description	edit a Group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//
//	@Param			group	body		groupservice.Group	true	"Group"
//
//	@Param			groupId	path		string				true	"Group ID"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Router			/groups/{groupId} [patch]
func (rest *Rest) editGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
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
	err = (*(*rest).gs).EditGroup(group, gID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
//
//	@Summary		delete a Group
//	@Description	delete a Group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			groupId	path		string	true	"Group ID"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Router			/groups/{groupId} [delete]
func (rest *Rest) deleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	gID, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = (*(*rest).gs).DeleteGroup(gID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
//
//	@Summary		add a user to a Group
//	@Description	add a user to a Group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			groupId	path		string	true	"Group ID"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Router			/groups/{groupId}/users [post]
func (rest *Rest) addUserToGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(r.Header.Get("userId"))
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
	err = (*(*rest).gs).AddUserToGroup(gID, uID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Swagger API Documentation
//
//	@Summary		remove a user from a Group
//	@Description	remove a user from a Group.
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			groupId	path		string	true	"Group ID"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Router			/groups/{groupId}/users/{userId} [delete]
func (rest *Rest) removeUserFromGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(r.Header.Get("userId"))

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
	err = (*(*rest).gs).RemoveUserFromGroup(gID, uID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
