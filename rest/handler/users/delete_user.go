package users

import (
	"backend-golang/repo"
	"backend-golang/response"
	"backend-golang/util"
	"errors"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	_, err = h.repo.Delete(id)
	if errors.Is(err, repo.ErrUserNotFound) {
		util.SendError(w, http.StatusNotFound, "User Not Found !")
		return
	}
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendDataWithMessage(w, "User deleted successfully", nil, http.StatusOK)
}
