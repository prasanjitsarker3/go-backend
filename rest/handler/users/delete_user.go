package users

import (
	"backend-golang/domain"
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

	_, err = h.service.Delete(id)
	if errors.Is(err, domain.ErrUserNotFound) {
		
		util.SendError(w, http.StatusNotFound, "User Not Found !")
		return
	}
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendDataWithMessage(w, "User deleted successfully", nil, http.StatusOK)
}
