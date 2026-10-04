package users

import (
	"backend-golang/repo"
	"backend-golang/response"
	"backend-golang/util"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	var user repo.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid requested body")
		return
	}
	user.ID = id

	updated, err := h.repo.Update(user)
	if errors.Is(err, repo.ErrUserNotFound) {
		util.SendError(w, http.StatusNotFound, "User Not Found !")
		return
	}
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendDataWithMessage(w, "User updated successfully", updated, http.StatusOK)
}
