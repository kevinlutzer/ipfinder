package rest

import (
	"encoding/json"
	"net/http"
)

type apiResponseType string

const (
	Success apiResponseType = "success"
	Failure apiResponseType = "error"
)

type ApiResponse[D any] struct {
	Type apiResponseType `json:"type" binding:"required"`
	Data D               `json:"data"`
}

func WriteSuccessResponse[D any](w http.ResponseWriter, data D) {
	w.Header().Set("Content-Type", "application/json")

	r := &ApiResponse[D]{
		Data: data,
		Type: Success,
	}

	b, _ := json.Marshal(r)
	w.WriteHeader(200)
	w.Write(b)

	return
}

func WriteErrorResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")

	r := &ApiResponse[string]{
		Data: msg,
		Type: Success,
	}

	b, _ := json.Marshal(r)
	w.WriteHeader(code)
	w.Write(b)
}
