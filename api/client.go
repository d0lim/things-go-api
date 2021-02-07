package api

import (
	"errors"
	"net/http"
)

const (
	// APIEndpoint is the public culturedcode https endpoint
	APIEndpoint = "https://cloud.culturedcode.com/version/1"
)

var (
	// ErrUnauthorized is returned by the API when the credentials are wrong
	ErrUnauthorized = errors.New("unauthorized")
)

// New initializes a things client
func New(endpoint, email, password string) *Client {
	c := &Client{
		Endpoint: endpoint,
		EMail:    email,
		password: password,

		client: &http.Client{},
	}
	c.commonSvc.client = c
	c.Accounts = (*AccountService)(&c.commonSvc)
	return c
}
