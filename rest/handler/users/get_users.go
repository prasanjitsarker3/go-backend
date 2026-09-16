package users

import (
	"backend-golang/response"
	"net/http"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response.SendData(w, h.repo.List(), http.StatusOK)
}
