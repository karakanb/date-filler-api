package main

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func (e Error) json() []byte {
	message, _ := json.Marshal(e)
	return message
}

func jsonError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(Error{message}.json())
}
