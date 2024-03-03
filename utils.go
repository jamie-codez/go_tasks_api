package main

import (
	"encoding/json"
	"net/http"
)

// Serializes the error message and sends it to the client
func sendError(res http.ResponseWriter, err error, statusCode int) {
	http.Error(res, err.Error(), statusCode)
}

// Serializes the response and sends it to the client
func sendResponse(res http.ResponseWriter, data interface{}, statusCode int) {
	res.WriteHeader(statusCode)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		sendError(res, err, http.StatusInternalServerError)
	}
}
