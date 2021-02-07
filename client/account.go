package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/d0lim/things-go-api/common"
)

// Delete deletes your current thingscloud account. This cannot be reversed
func (s *common.AccountService) Delete() error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/account/%s", s.client.EMail), nil)
	if err != nil {
		return err
	}
	resp, err := s.client.do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		if resp.StatusCode == http.StatusUnauthorized {
			return ErrUnauthorized
		}
		return fmt.Errorf("http response code: %s", resp.Status)
	}
	return nil
}

// Confirm finishes the account creation by providing the email token send by thingscloud
func (s *common.AccountService) Confirm(code string) error {
	data, err := json.Marshal(common.accountRequestBody{
		ConfirmationCode: code,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("/account/%s", s.client.EMail), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := s.client.do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return ErrUnauthorized
		}
		return fmt.Errorf("http response code: %s", resp.Status)
	}
	return nil
}

// SignUp creates a new thingscloud account and returns a configured client
func (s *common.AccountService) SignUp(email, password string) (*common.Client, error) {
	data, err := json.Marshal(common.accountRequestBody{
		Password:           password,
		SLAVersionAccepted: "https://thingscloud.appspot.com/sla/v5.html",
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("/account/%s", email), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	resp, err := s.client.do(req)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("http response code: %s", resp.Status)
	}

	return New(s.client.Endpoint, email, password), nil
}

// ChangePassword allows you to change your account password.
// Because things does not work with sessions you need to create a new client instance after
// executing this method
func (s *common.AccountService) ChangePassword(newPassword string) (*common.Client, error) {
	data, err := json.Marshal(common.accountRequestBody{
		Password: newPassword,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("/account/%s", s.client.EMail), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	resp, err := s.client.do(req)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, ErrUnauthorized
		}
		return nil, fmt.Errorf("http response code: %s", resp.Status)
	}

	return New(s.client.Endpoint, s.client.EMail, newPassword), nil
}
