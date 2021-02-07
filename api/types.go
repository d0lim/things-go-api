package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ThingsUserAgent is the http user-agent header set by things for mac Version 3.1.0(30100506)
const ThingsUserAgent = "ThingsMac/30100506mas"

type accountRequestBody struct {
	Password           string `json:"password,omitempty"`
	SLAVersionAccepted string `json:"SLA-version-accepted,omitempty"`
	ConfirmationCode   string `json:"confirmation-code,omitempty"`
}

// AccountService allows account specific interaction with thingscloud
type AccountService service

// Client is a culturedcode cloud client. It can be used to interact with the
// things cloud to manage your data.
type Client struct {
	Endpoint string
	EMail    string
	password string

	client    *http.Client
	commonSvc service

	Accounts *AccountService
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
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

type service struct {
	client *Client
}

// TaskStatus describes if a thing is completed or not
type TaskStatus int

const (
	// TaskStatusPending indicates a new task
	TaskStatusPending TaskStatus = 0
	// TaskStatusCompleted indicates a completed task
	TaskStatusCompleted TaskStatus = 3
	// TaskStatusCanceled indicates a canceled task
	TaskStatusCanceled TaskStatus = 2
)

// TaskSchedule describes when a task is scheduled
type TaskSchedule int

const (
	// TaskScheduleToday indicates tasks which should be completed today
	TaskScheduleToday TaskSchedule = 0
	// TaskScheduleAnytime indicates tasks which can be completed anyday
	TaskScheduleAnytime TaskSchedule = 1
	// TaskScheduleSomeday indicates tasks which might never be completed
	TaskScheduleSomeday TaskSchedule = 2
)

// Timestamp allows unix epochs represented as float or ints to be unmarshalled
// into time.Time objects
type Timestamp time.Time

// UnmarshalJSON takes a unix epoch from float/ int and creates a time.Time instance
func (t *Timestamp) UnmarshalJSON(bs []byte) error {
	var d float64
	if err := json.Unmarshal(bs, &d); err != nil {
		return err
	}
	*t = Timestamp(time.Unix(int64(d), 0).UTC())
	return nil
}

// MarshalJSON convers a timestamp into unix nano representation
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	var tt = time.Time(*t).Unix()
	return json.Marshal(tt)
}

// Format returns a textual representation of the time value formatted according to layout
func (t *Timestamp) Format(layout string) string {
	return time.Time(*t).Format(layout)
}

// Time returns the underlying time.Time instance
func (t *Timestamp) Time() *time.Time {
	tt := time.Time(*t)
	return &tt
}
