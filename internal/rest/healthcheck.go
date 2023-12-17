package rest

import "net/http"

func (r *rest) HealthCheckHander(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet && req.Method != http.MethodOptions {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Only GET and OPTION methods are allowed")
		return
	}

	WriteSuccessResponse[string](w, "Alive")
}
