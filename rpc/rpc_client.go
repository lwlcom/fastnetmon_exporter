package rpc

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"log"
)

// Client sends commands the Fastnetmon API
type Client struct {
	Debug       bool
	apiHost     string
	apiUsername string
	apiPassword string
}

// NewClient creates a new client connection
func NewClient(apiHost, apiUsername, apiPassword string, debug bool) *Client {
	rpc := &Client{debug, apiHost, apiUsername, apiPassword}

	return rpc
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// RunCommand requests a path from the API
func (c *Client) RunCommand(path string) ([]byte, error) {
	var client http.Client

	host := c.apiHost
	if !strings.HasPrefix("host", "http") {
		host = "http://" + host
	}

	req, err := http.NewRequest("GET", host+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(c.apiUsername, c.apiPassword))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// RunCommandAndParse makes an API request and parses the JSON result
func (c *Client) RunCommandAndParse(path string, obj interface{}) error {
	if c.Debug {
		log.Printf("Requesting: %s\n", path)
	}

	b, err := c.RunCommand(path)
	if err != nil {
		return err
	}

	if c.Debug {
		log.Printf("Output for %s: %s\n", c.apiHost, string(b))
	}

	err = json.Unmarshal(b, obj)
	return err
}
