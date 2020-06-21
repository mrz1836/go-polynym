package polynym

import (
	"fmt"
	"net/http"
	"time"

	"testing"
)

// newMockClient will create a new mock client for testing
func newMockClient(userAgent string) Client {
	return Client{
		httpClient: &mockHTTP{},
		UserAgent:  userAgent,
	}
}

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	client := NewClient(nil)

	if len(client.UserAgent) == 0 {
		t.Fatal("missing user agent")
	}
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client := NewClient(nil)
	fmt.Println(client.UserAgent)
	// Output:go-polynym: v0.3.0
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewClient(nil)
	}
}

// TestClientDefaultOptions tests setting ClientDefaultOptions()
func TestClientDefaultOptions(t *testing.T) {

	options := ClientDefaultOptions()

	if options.UserAgent != defaultUserAgent {
		t.Fatalf("expected value: %s got: %s", defaultUserAgent, options.UserAgent)
	}

	if options.BackOffExponentFactor != 2.0 {
		t.Fatalf("expected value: %f got: %f", 2.0, options.BackOffExponentFactor)
	}

	if options.BackOffInitialTimeout != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffInitialTimeout)
	}

	if options.BackOffMaximumJitterInterval != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffMaximumJitterInterval)
	}

	if options.BackOffMaxTimeout != 10*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 10*time.Millisecond, options.BackOffMaxTimeout)
	}

	if options.DialerKeepAlive != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.DialerKeepAlive)
	}

	if options.DialerTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.DialerTimeout)
	}

	if options.RequestRetryCount != 2 {
		t.Fatalf("expected value: %v got: %v", 2, options.RequestRetryCount)
	}

	if options.RequestTimeout != 10*time.Second {
		t.Fatalf("expected value: %v got: %v", 10*time.Second, options.RequestTimeout)
	}

	if options.TransportExpectContinueTimeout != 3*time.Second {
		t.Fatalf("expected value: %v got: %v", 3*time.Second, options.TransportExpectContinueTimeout)
	}

	if options.TransportIdleTimeout != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.TransportIdleTimeout)
	}

	if options.TransportMaxIdleConnections != 10 {
		t.Fatalf("expected value: %v got: %v", 10, options.TransportMaxIdleConnections)
	}

	if options.TransportTLSHandshakeTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.TransportTLSHandshakeTimeout)
	}
}

// TestClientDefaultOptions_NoRetry will set 0 retry counts
func TestClientDefaultOptions_NoRetry(t *testing.T) {
	options := ClientDefaultOptions()
	options.RequestRetryCount = 0
	client := NewClient(options)

	if client.UserAgent != defaultUserAgent {
		t.Errorf("user agent mismatch")
	}
}

// TestGetAddress tests the GetAddress()
func TestGetAddress(t *testing.T) {

	// Create a mock client
	client := newMockClient(defaultUserAgent)

	// Create the list of tests
	var tests = []struct {
		input         string
		expected      string
		expectedError bool
		statusCode    int
	}{
		{"", "", true, http.StatusBadRequest},
		{"error", "", true, http.StatusBadRequest},
		{"bad-poly-response", "", true, http.StatusBadRequest},
		{"bad-poly-status", "", true, http.StatusBadRequest},
		{"doesnotexist@handcash.io", "", true, http.StatusBadRequest},
		{"$mr-z", "124dwBFyFtkcNXGfVWQroGcT9ybnpQ3G3Z", false, http.StatusOK},
		{"19gKzz8XmFDyrpk4qFobG7qKoqybe78v9h", "19gKzz8XmFDyrpk4qFobG7qKoqybe78v9h", false, http.StatusOK},
		{"1doesnotexisthandle", "", true, http.StatusBadRequest},
		{"1mrz", "1Lti3s6AQNKTSgxnTyBREMa6XdHLBnPSKa", false, http.StatusOK},
		{"bad@paymailaddress.com", "", true, http.StatusBadRequest},
		{"c6ZqP5Tb22KJuvSAbjNkoi", "", true, http.StatusBadRequest},
		{"mrz@handcash.io", "19gKzz8XmFDyrpk4qFobG7qKoqybe78v9h", false, http.StatusOK},
		{"@833", "19ksW6ueSw9nEj88X3QNJ9VkKPGf1zuKbQ", false, http.StatusOK},
	}

	// Test all
	for _, test := range tests {
		if output, err := GetAddress(client, test.input); err == nil && test.expectedError {
			t.Errorf("%s Failed: expected to throw an error, no error [%s] inputted and [%s] expected", t.Name(), test.input, test.expected)
		} else if err != nil && !test.expectedError {
			t.Errorf("%s Failed: [%s] inputted and [%s] expected, received: [%v] error [%s]", t.Name(), test.input, test.expected, output, err.Error())
		} else if output != nil && output.Address != test.expected && !test.expectedError {
			t.Errorf("%s Failed: [%s] inputted and [%s] expected, received: [%s]", t.Name(), test.input, test.expected, output.Address)
		} else if output != nil && output.LastRequest.Method != http.MethodGet {
			t.Errorf("%s Expected method to be %s, got %s, [%s] inputted", t.Name(), http.MethodGet, output.LastRequest.Method, test.input)
		} else if output != nil && output.LastRequest.StatusCode != test.statusCode {
			t.Errorf("%s Expected status code to be %d, got %d, [%s] inputted", t.Name(), test.statusCode, output.LastRequest.StatusCode, test.input)
		}
	}
}

// ExampleGetAddress example using GetAddress()
func ExampleGetAddress() {
	client := newMockClient(defaultUserAgent)
	resp, _ := GetAddress(client, "16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA")
	fmt.Println(resp.Address)
	// Output:16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA
}

// TestHandCashConvert will test the HandCashConvert() method
func TestHandCashConvert(t *testing.T) {
	// Create the list of tests
	var tests = []struct {
		input    string
		expected string
	}{
		{"$mr-z", "mr-z@handcash.io"},
		{"invalid$mr-z", "invalid$mr-z"},
		{"$", "@handcash.io"},
		{"1handle", "1handle"},
	}

	// Test all
	for _, test := range tests {
		if output := HandCashConvert(test.input); output != test.expected {
			t.Errorf("%s Failed: [%s] inputted and [%s] expected, received: [%s]", t.Name(), test.input, test.expected, output)
		}
	}
}

// ExampleHandCashConvert example using HandCashConvert()
func ExampleHandCashConvert() {
	paymail := HandCashConvert("$mr-z")
	fmt.Println(paymail)
	// Output:mr-z@handcash.io
}

// BenchmarkHandCashConvert benchmarks the HandCashConvert method
func BenchmarkHandCashConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HandCashConvert("$mr-z")
	}
}
