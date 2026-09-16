package product

import (
	"backend-golang/response"
	"net/http"
)

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	response.SendData(w, h.repo.List(), http.StatusOK)
}
