package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/d0lim/things-go-api/common"
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
func New(endpoint, email, password string) *common.Client {
	c := &common.Client{
		Endpoint: endpoint,
		EMail:    email,
		password: password,

		client: &http.Client{},
	}
	c.common.client = c
	c.Accounts = (*common.AccountService)(&c.common)
	return c
}

// ThingsUserAgent is the http user-agent header set by things for mac Version 3.1.0(30100506)
const ThingsUserAgent = "ThingsMac/30100506mas"

func (c *common.Client) do(req *http.Request) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s", c.Endpoint, req.URL)
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	req.URL = u

	req.Header.Set("Authorization", fmt.Sprintf("Password %s", c.password))
	req.Header.Set("User-Agent", ThingsUserAgent)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Encoding", "UTF8")
	req.Header.Set("Accept-Language", "en-us")

	return c.client.Do(req)
}
