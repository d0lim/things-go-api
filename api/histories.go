package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// History represents a synchronization stream. It's identified with a uuid v4
type History struct {
	ID                     string
	Client                 *Client
	LatestServerIndex      int
	LoadedServerIndex      int
	LatestSchemaVersion    int
	EndTotalContentSize    int
	LatestTotalContentSize int
}

type historyResponse struct {
	LatestSchemaVersion    int  `json:"latest-schema-version"`
	LatestTotalContentSize int  `json:"latest-total-content-size"`
	IsEmpty                bool `json:"is-empty"`
	LatestServerIndex      int  `json:"latest-server-index"`
}

// Sync ensures the history object is able to write to things
func (h *History) Sync() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("/history/%s", h.ID), nil)
	if err != nil {
		return err
	}
	resp, err := h.Client.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http response code: %s", resp.Status)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var v historyResponse
	json.Unmarshal(bs, &v)
	h.LatestServerIndex = v.LatestServerIndex
	h.LatestSchemaVersion = v.LatestSchemaVersion
	h.LatestTotalContentSize = v.LatestTotalContentSize
	return nil
}

// History requests a specific history
func (c *Client) History(id string) (*History, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/history/%s", id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, ErrUnauthorized
		}
		return nil, fmt.Errorf("http response code: %s", resp.Status)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	h := historyResponse{}
	if err := json.Unmarshal(bs, &h); err != nil {
		return nil, err
	}

	return &History{
		Client:                 c,
		ID:                     id,
		LatestServerIndex:      h.LatestServerIndex,
		LatestSchemaVersion:    h.LatestSchemaVersion,
		LatestTotalContentSize: h.LatestTotalContentSize,
	}, nil
}

type v1historyResponse struct {
	Key                 string `json:"history-key"`
	LatestServerIndex   int    `json:"latest-server-index"`
	IsEmpty             bool   `json:"is-empty"`
	LatestSchemaVersion int    `json:"latest-schema-version"`
}

// OwnHistory returns the clients own history
func (c *Client) OwnHistory() (*History, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/account/%s/own-history-key", c.EMail), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, ErrUnauthorized
		}
		return nil, fmt.Errorf("http response code: %s", resp.Status)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data v1historyResponse
	json.Unmarshal(bs, &data)

	return &History{
		Client: c,
		ID:     data.Key,
	}, nil
}
