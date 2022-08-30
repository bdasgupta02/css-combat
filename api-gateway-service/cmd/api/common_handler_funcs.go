package main

import (
	"net/http"
	"encoding/json"
)

func writeJSONResponse(w http.ResponseWriter, jsonData any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func (app *serverConfig) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	writeJSONResponse(w, payload)
}


type jsonResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}