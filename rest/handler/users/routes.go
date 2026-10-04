package users

import (
	"net/http"
)

func (h *Handler) UserRegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", h.CreateNewUser)
	mux.HandleFunc("GET /api/users", h.GetAllUsers)
	mux.HandleFunc("GET /api/users/{id}", h.GetSingleUser)
	mux.HandleFunc("PUT /api/users/{id}", h.UpdateUser)
	mux.HandleFunc("DELETE /api/users/{id}", h.DeleteUser)
}
