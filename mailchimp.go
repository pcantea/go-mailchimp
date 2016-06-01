package mailchimp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client manages communication with the Mailchimp API.
type Client struct {
	client  *http.Client
	BaseURL *url.URL
	DC      string
	APIKey  string
}

// NewClient returns a new Mailchimp API client.  If a nil httpClient is
// provided, http.DefaultClient will be used. The apiKey must be in the format xyz-us11.
func NewClient(apiKey string, httpClient *http.Client) (*Client, error) {
	if len(strings.Split(apiKey, "-")) != 2 {
		return nil, errors.New("Mailchimp API Key must be formatted like: xyz-zys")
	}
	dc := strings.Split(apiKey, "-")[1] // data center
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(fmt.Sprintf("https://%s.api.mailchimp.com/3.0", dc))
	return &Client{
		APIKey:  apiKey,
		client:  httpClient,
		DC:      dc,
		BaseURL: baseURL,
	}, nil
}

// ErrorResponse ...
type ErrorResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

// Error ...
func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d %s (%s)", e.Status, e.Title, e.Detail)
}

// CheckResponse ...
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := new(ErrorResponse)
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// Do ...
func (c *Client) Do(method string, path string, body interface{}) (interface{}, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	apiURL := fmt.Sprintf("%s%s", c.BaseURL.String(), path)

	req, err := http.NewRequest(method, apiURL, buf)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("", c.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	var v interface{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Subscribe ...
func (c *Client) Subscribe(email string, listID string) (interface{}, error) {
	data := &map[string]string{
		"email_address": email,
		"status":        "subscribed",
	}
	return c.Do(
		"POST",
		fmt.Sprintf("/lists/%s/members/", listID),
		data,
	)
}
