package users

import (
	"backend-golang/response"
	"backend-golang/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetSingleUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	user, err := h.repo.Get(id)
	if err != nil {
		util.SendError(w, http.StatusNotFound, "User Not Found !")
		return
	}

	response.SendDataWithMessage(w, "User retrieved successfully", user, http.StatusOK)
}
