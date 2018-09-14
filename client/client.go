package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	ENDPOINT_URL     = "https://api.authorize.net/xml/v1/request.api"     // Authorize.Net API endpoint
	DEV_ENDPOINT_URL = "https://apitest.authorize.net/xml/v1/request.api" // Authorize.Net API DEV endpoint
)

type Client struct {
	url     string
	sandbox bool
}

// NewClient creates a new Authorize API client instance.
func NewClient(sandbox bool) *Client {
	url := ENDPOINT_URL
	if sandbox {
		url = DEV_ENDPOINT_URL
	}

	return &Client{
		url:     url,
		sandbox: sandbox,
	}
}

func (c *Client) PostRequest(bodyObject interface{}) (output []byte, err error) {
	req, err := c.createRequest("POST", bodyObject)
	if err != nil {
		return nil, err
	}

	return c.executeRequest(req)
}

func (c *Client) createRequest(method string, bodyObject interface{}) (req *http.Request, err error) {
	var reqBody io.Reader
	if bodyObject != nil {
		data, err := xml.Marshal(bodyObject)
		if err != nil {
			return nil, fmt.Errorf("Error marshaling body object: %s", err.Error())
		}

		reqBody = bytes.NewBuffer(data)
	}

	req, err = http.NewRequest(method, c.url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")

	// no keep-alive
	req.Header.Set("Connection", "close")
	req.Close = true

	return req, nil
}

func (c *Client) executeRequest(req *http.Request) (output []byte, err error) {
	httpClient := http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making HTTP request: %s", err.Error())
	}
	defer res.Body.Close()

	output, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body data: %s", err.Error())
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("Error %s: %s", res.Status, res.StatusCode)
	}

	return output, nil
}
