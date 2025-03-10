package http_controller

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func ValidationErrorRespond(w http.ResponseWriter, validationError *ValidationError) {
	Respond(w, http.StatusBadRequest, validationError)
}

func InternalServerErrorRespond(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func NotFoundErrorRespond(w http.ResponseWriter, notFoundError *NotFoundError) {
	Respond(w, http.StatusNotFound, notFoundError)
}
