package main

import (
	"encoding/json"
	"net/http"
)

// Serializes the response and sends it to the client
func sendResponse(res http.ResponseWriter, statusCode int, statusMessage string, hasErrors bool, message string, data interface{}) {
	res.WriteHeader(statusCode)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewEncoder(res).Encode(Response{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		HasErrors:     hasErrors,
		Message:       message,
		Data:          data,
	})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		_ = json.NewEncoder(res).Encode(Response{
			StatusCode:    http.StatusInternalServerError,
			StatusMessage: "Internal Server Error",
			HasErrors:     true,
			Message:       "Failed to serialize response",
			Data:          nil,
		})
	}
}
