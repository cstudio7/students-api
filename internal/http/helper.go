package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string
	Error   string
}

func respondWithError(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
