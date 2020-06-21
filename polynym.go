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
	Address      string       `json:"address"`
	ErrorMessage string       `json:"error"`
	LastRequest  *LastRequest `json:"-"`
}

// GetAddress returns the address of a given 1handle, $handcash, paymail, Twetch user id or BitcoinSV address
func GetAddress(client Client, handleOrAddress string) (response *GetAddressResponse, err error) {

	// Set the API url
	reqURL := fmt.Sprintf("%s/%s/%s", apiEndpoint, "getAddress", HandCashConvert(handleOrAddress))

	// Store for debugging purposes
	response = &GetAddressResponse{
		LastRequest: &LastRequest{
			Method:   http.MethodGet,
			PostData: reqURL,
			URL:      reqURL,
		},
	}

	// Check for a value
	if len(handleOrAddress) == 0 {
		response.LastRequest.StatusCode = http.StatusBadRequest
		err = fmt.Errorf("missing handle or paymail to resolve")
		return
	}

	// Start the request
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, reqURL, nil); err != nil {
		return
	}

	// Set the header (user agent is in case they block default Go user agents)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.UserAgent)

	// Fire the request
	var resp *http.Response
	if resp, err = client.httpClient.Do(req); err != nil {
		return
	}

	// Cleanup
	defer func() {
		_ = resp.Body.Close()
	}()

	// Set the status
	response.LastRequest.StatusCode = resp.StatusCode

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

	// Try and decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)

	return
}

// HandCashConvert now converts $handles to handle@handcash.io
func HandCashConvert(handle string) string {
	if strings.HasPrefix(handle, "$") {
		handle = strings.Replace(handle, "$", "", -1) + "@handcash.io"
	}
	return handle
}
