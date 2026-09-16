package users

import (
	"backend-golang/repo"
	"backend-golang/response"
	"backend-golang/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user repo.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid requested body")
		return
	}

	created, err := h.repo.Create(user)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendData(w, created, http.StatusCreated)
}
