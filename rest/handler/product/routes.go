package product

import (
	"net/http"
)

func (h *Handler) ProductRegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/products", h.CreateProduct)
	mux.HandleFunc("GET /api/products", h.GetAllProducts)
	mux.HandleFunc("GET /api/products/{id}", h.GetSingleProduct)
}
