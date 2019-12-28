/*
Package polynym is the unofficial golang implementation for the Polynym API

Example:

// Create a new client
client, _ := polynym.NewClient(nil)

// Get address
resp, _ := client.GetAddress("1mrz")

// What's the address?
log.Println("address:", resp.Address)
*/
package polynym

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAddressResponse is what polynym returns (success or fail)
type GetAddressResponse struct {
	Address      string `json:"address"`
	ErrorMessage string `json:"error"`
}

// NewClient creates a new client to submit requests
func NewClient(clientOptions *Options) (c *Client, err error) {

	// Create a client using the given options
	c = createClient(clientOptions)

	return
}

// GetAddress returns the address of a given 1handle, $handcash, paymail or BitcoinSV address
func (c *Client) GetAddress(handleOrAddress string) (response *GetAddressResponse, err error) {

	// Set the API url
	reqURL := fmt.Sprintf("%s/%s/%s", apiEndpoint, "getAddress", detectHandCash(handleOrAddress))

	// Store for debugging purposes
	c.LastRequest.Method = http.MethodGet
	c.LastRequest.URL = reqURL
	c.LastRequest.PostData = reqURL

	// Start the request
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, reqURL, nil); err != nil {
		return
	}

	// Set the header (user agent is in case they block default Go user agents)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.Parameters.UserAgent)

	// Fire the request
	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	// Cleanup
	defer func() {
		_ = resp.Body.Close()
	}()

	// Save the status
	c.LastRequest.StatusCode = resp.StatusCode

	// Handle errors
	if resp.StatusCode != http.StatusOK {

		// Decode the error message
		if resp.StatusCode == http.StatusBadRequest {
			if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
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

// detectHandCash now converts $handles to handle@handcash.io
func detectHandCash(handle string) string {
	if strings.Contains(handle, "$") {
		handle = strings.Replace(handle, "$", "", -1) + "@handcash.io"
	}
	return handle
}
