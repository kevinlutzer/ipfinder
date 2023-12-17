package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	clientIPAPI = "api.bigdatacloud.net/data/client-ip"
)

type IPService interface {
	GetPublicIP() (string, error)
}

type ipservice struct{}

func NewIPService() IPService {
	return &ipservice{}
}

type apiResp struct {
	IPAddress string `json:"ipString"`
}

func (s *ipservice) GetPublicIP() (string, error) {
	resp, err := http.Get("https://" + clientIPAPI)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Resp from internal API was a non 200 code: %s", resp.Status))
	}

	// Load response body
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarhsal API response
	d := &apiResp{}
	if err := json.Unmarshal(b, d); err != nil {
		return "", err
	}

	return d.IPAddress, nil
}
