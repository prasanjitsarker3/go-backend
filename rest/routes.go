package rest

import (
	"backend-golang/response"
	"net/http"
)

func routers ( mux *http.ServeMux ) {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		response.SendData(w, map[string]string{"message": "Server running successfully!"}, http.StatusOK)
	})
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		response.SendData(w, map[string]string{"status": "healthy"}, http.StatusOK)
	})
}