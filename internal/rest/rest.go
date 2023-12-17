package rest

import (
	"github.com/gorilla/mux"
	"github.com/kevinlutzer/ip-finder/internal/service"
)

type rest struct {
	ipService service.IPService
	apiKey    string
}

func NewRest(ipService service.IPService, apiKey string) *mux.Router {
	m := mux.NewRouter()

	r := &rest{
		ipService: ipService,
		apiKey:    apiKey,
	}

	m.HandleFunc("/v1/publicip", r.PublicIPHandler)
	m.HandleFunc("/v1/healthcheck", r.HealthCheckHander)

	return m
}
