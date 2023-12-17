package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/kevinlutzer/ip-finder/internal/rest"
	"github.com/kevinlutzer/ip-finder/internal/service"
)

const (
	ErrServerClosedCode = 4
	ErrServerClosed     = 5
	ErrAPIKeyIsRequired = 6
)

func main() {

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		os.Exit(ErrAPIKeyIsRequired)
	}

	s := service.NewIPService()
	m := rest.NewRest(s, apiKey)

	if err := http.ListenAndServe(":80", m); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			os.Exit(ErrServerClosedCode)
		}

		os.Exit(ErrServerClosed)
	}
}
