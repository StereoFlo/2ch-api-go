package utils

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

func Respond(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getMessage(true, data))
}

func RespondError(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(getMessage(false, data))
}

func getMessage(success bool, data interface{}) Message {
	return Message{
		Success: success,
		Message: data,
	}
}
