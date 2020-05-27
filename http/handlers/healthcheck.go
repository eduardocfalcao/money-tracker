package handlers

import (
	"encoding/json"
	"net/http"
)

type status struct {
	Health string `json:"health"`
}

// Healthcheck is a handler to return the status of the service
func Healthcheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	var status = status{
		Health: "alive",
	}

	_ = json.NewEncoder(w).Encode(status)
}