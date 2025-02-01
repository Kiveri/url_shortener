package controller

import (
	"encoding/json"
	"net/http"
)

func Validation(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func ValidationErrorRespond(w http.ResponseWriter, validationError *ValidationError) {
	Validation(w, http.StatusBadRequest, validationError)
}

func InternalServer(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
