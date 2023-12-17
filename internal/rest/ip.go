package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PublicIPAPIResp struct {
	IPAddress string `json: "ipAddress"`
}

type PublicIPAPIReq struct {
	APIKey string `json:"apiKey"`
}

func (r *rest) PublicIPHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost && req.Method != http.MethodOptions {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Only the POST method is allowed for this API")
		return
	}

	// Read body
	b, err := io.ReadAll(req.Body)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Request body is not valid JSON")
		return
	}

	// Unmarshal Request
	d := &PublicIPAPIReq{}
	if err := json.Unmarshal(b, d); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Request couldn't be unmarshalled")
		return
	}

	// Validate Request
	if d.APIKey == "" {
		WriteErrorResponse(w, http.StatusBadRequest, "apiKey must be present in the request")
		return
	}

	// Validate Auth
	if r.apiKey != d.APIKey {
		WriteErrorResponse(w, http.StatusUnauthorized, "Unauthorized request")
		return
	}

	// Get the Public IP
	ipAddress, err := r.ipService.GetPublicIP()
	if err != nil {
		msg := fmt.Sprintf("Failed to get the public IP with error: %s", err.Error())
		WriteErrorResponse(w, http.StatusInternalServerError, msg)
	}

	// Success!
	WriteSuccessResponse[PublicIPAPIResp](w, PublicIPAPIResp{
		IPAddress: ipAddress,
	})
}
