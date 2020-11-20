package Server

import (
	"encoding/json"
	"net/http"
)

type ResponseFormat struct {
	Message string      `json: "message"`
	Detail  interface{} `json: "detail"`
}

type IResponseFormat interface {
	FormatResponse(w http.ResponseWriter, status int)
}

func NewResponseFormat(message string, detail interface{}) IResponseFormat {
	return &ResponseFormat{
		Message: message,
		Detail:  detail,
	}
}

func (f *ResponseFormat) FormatResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "JSON")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(f)
	if err != nil {
		resp := NewResponseFormat("Internal Server Error", nil)
		resp.FormatResponse(w, http.StatusInternalServerError)
	}
}
