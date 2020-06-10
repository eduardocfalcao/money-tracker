package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type status struct {
	Health string `json:"health"`
	Host   string `json:host`
}

// Healthcheck is a handler to return the status of the service
func Healthcheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	host, _ := os.Hostname()
	var status = status{
		Health: "alive",
		Host:   host,
	}

	_ = json.NewEncoder(w).Encode(status)
}
