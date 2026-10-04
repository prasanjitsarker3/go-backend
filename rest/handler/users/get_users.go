package users

import (
	"backend-golang/response"
	"net/http"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response.SendDataWithMessage(w, "Users retrieved successfully", h.repo.List(), http.StatusOK)
}
