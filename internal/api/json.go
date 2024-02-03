package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
	JsonResponseWithCode(http.StatusOK, w, data)
}

func JsonResponseWithCode(statusCode int, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		InternalErrorResponse(w)
		logrus.Errorf("Error serializing object to the response: %s", err)
		return
	}
	w.WriteHeader(statusCode)
}
