/*
Package polynym is the unofficial golang implementation for the Polynym API

Example:

// Create a new client
client, _ := polynym.NewClient()

// Get address
resp, _ := client.GetAddress("1mrz")

// What's the address?
log.Println("address:", resp.Address)
*/
package polynym

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

// Client holds client configuration settings
type Client struct {

	// HTTPClient carries out the POST operations
	HTTPClient heimdall.Client

	// LastRequest is the raw information from the last request
	LastRequest *LastRequest

	// UserAgent (optional for changing user agents)
	UserAgent string
}

// LastRequest is used to track what was submitted to the Request()
type LastRequest struct {

	// Method is either POST or GET
	Method string

	// PostData is the post data submitted if POST request
	PostData string

	// StatusCode is the last code from the request
	StatusCode int

	// URL is the url used for the request
	URL string
}

// NewClient creates a new client to submit requests
func NewClient() (c *Client, err error) {

	// Create a client
	c = new(Client)

	// Create exponential backoff
	backOff := heimdall.NewExponentialBackoff(
		ConnectionInitialTimeout,
		ConnectionMaxTimeout,
		ConnectionExponentFactor,
		ConnectionMaximumJitterInterval,
	)

	// Create the http client
	c.HTTPClient = httpclient.NewClient(
		httpclient.WithHTTPTimeout(ConnectionWithHTTPTimeout),
		httpclient.WithRetrier(heimdall.NewRetrier(backOff)),
		httpclient.WithRetryCount(ConnectionRetryCount),
		httpclient.WithHTTPClient(&http.Client{
			Transport: ClientDefaultTransport,
		}),
	)

	// Set defaults
	c.UserAgent = DefaultUserAgent

	// Create a last request struct
	c.LastRequest = new(LastRequest)

	// Return the client
	return
}

// GetAddressResponse is what polynym returns (success or fail)
type GetAddressResponse struct {
	Address      string `json:"address"`
	ErrorMessage string `json:"error"`
}

// GetAddress returns the address of a given 1handle, paymail or BSV address ($handcash deprecated)
func (c *Client) GetAddress(handleOrAddress string) (response *GetAddressResponse, err error) {

	// Set the API url
	reqURL := fmt.Sprintf("%s/%s/%s", APIEndpoint, "getAddress", handleOrAddress)

	// Store for debugging purposes
	c.LastRequest.Method = http.MethodGet
	c.LastRequest.URL = reqURL

	// Start the request
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, reqURL, nil); err != nil {
		return
	}

	// Set the header (user agent is in case they block default Go user agents)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	// Fire the request
	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	// Cleanup
	defer func() {
		if bodyErr := resp.Body.Close(); bodyErr != nil {
			log.Printf("error closing response body: %s", bodyErr.Error())
		}
	}()

	// Save the status
	c.LastRequest.StatusCode = resp.StatusCode

	// Handle errors
	if resp.StatusCode != 200 {

		// Decode the error message
		if resp.StatusCode == 400 {
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				return
			}
			if len(response.ErrorMessage) == 0 {
				response.ErrorMessage = "unknown error resolving address"
			}
			err = fmt.Errorf("error: %s", response.ErrorMessage)
		} else {
			err = fmt.Errorf("bad response from polynym: %d", resp.StatusCode)
		}

		return
	}

	// Try and decode the transactions
	err = json.NewDecoder(resp.Body).Decode(&response)

	return
}
