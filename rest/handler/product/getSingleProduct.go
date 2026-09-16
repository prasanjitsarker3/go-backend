package product

import (
	"backend-golang/response"
	"backend-golang/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	product, err := h.repo.Get(id)
	if err != nil {
		util.SendError(w, http.StatusNotFound, "Product Not Found !")
		return
	}

	response.SendData(w, product, http.StatusOK)
}
