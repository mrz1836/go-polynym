package polynym

import (
	"fmt"
	"time"

	"testing"
)

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	client, err := NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(client.Parameters.UserAgent) == 0 {
		t.Fatal("missing user agent")
	}
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient(nil)
	fmt.Println(client.Parameters.UserAgent)
	// Output:go-polynym: v1
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient(nil)
	}
}

// TestDefaultOptions tests setting ClientDefaultOptions()
func TestDefaultOptions(t *testing.T) {

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

// TestClient_GetAddress tests the GetAddress()
func TestClient_GetAddress(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries
	client, err := NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	address := "16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA"
	var resp *GetAddressResponse
	resp, err = client.GetAddress(address)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Address != address {
		t.Fatal("address should have resolved:", address)
	}

}

// TestClient_GetAddressRelayX tests the GetAddress()
func TestClient_GetAddressRelayX(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries
	client, err := NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	address := "1mrz"
	var resp *GetAddressResponse
	resp, err = client.GetAddress(address)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.Address) == 0 {
		t.Fatal("address should have resolved:", address)
	}

}

// TestClient_GetAddressPaymail tests the GetAddress()
func TestClient_GetAddressPaymail(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries
	client, err := NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	address := "mrz@moneybutton.com"
	var resp *GetAddressResponse
	resp, err = client.GetAddress(address)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.Address) == 0 {
		t.Fatal("address should have resolved:", address)
	}

}

// TestClient_GetAddressHandCash tests the GetAddressHandCash()
func TestClient_GetAddressHandCash(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries
	client, err := NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	address := "$handcash"
	var resp *GetAddressResponse
	resp, err = client.GetAddress(address)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.Address) == 0 {
		t.Fatal("address should have resolved:", address)
	}

}

// ExampleClient_GetAddress example using GetAddress()
func ExampleClient_GetAddress() {
	client, _ := NewClient(nil)
	resp, _ := client.GetAddress("16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA") //mrz@moneybutton.com
	fmt.Println(resp.Address)
	// Output:16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA
}
