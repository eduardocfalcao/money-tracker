package api

import (
	"net/http"
)

func UnauthorizedResponse(w http.ResponseWriter) {
	WriteApiError(w, http.StatusUnauthorized, APIError{
		Message: "Unauthorized request.",
	})
}

func MalformedJsonResponse(w http.ResponseWriter) {
	BadRequestResponse(w, "Malformed json received in the body.")
}

func BadRequestResponse(w http.ResponseWriter, message string) {
	WriteApiError(w, http.StatusBadRequest, APIError{Message: message})
}

func InternalErrorResponse(w http.ResponseWriter) {
	WriteApiError(w, http.StatusBadRequest, APIError{Message: "An internal error occurred."})
}

func WriteApiError(w http.ResponseWriter, httpStatus int, e APIError) {
	JsonResponseWithCode(httpStatus, w, e)
}
