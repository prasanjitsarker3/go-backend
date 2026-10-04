package response

import (
	"encoding/json"
	"net/http"
)

// APIResponse wraps a payload with a human-readable message.
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// SendDataWithMessage writes {"message": ..., "data": ...} as JSON.
func SendDataWithMessage(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	SendData(w, APIResponse{Message: message, Data: data}, statusCode)
}
