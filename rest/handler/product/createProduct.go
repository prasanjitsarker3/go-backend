package product

import (
	"backend-golang/repo"
	"backend-golang/response"
	"backend-golang/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product repo.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	created, err := h.repo.Create(product)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendData(w, created, http.StatusCreated)
}
