package polynym

import (
	"fmt"

	"testing"
)

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	if len(client.UserAgent) == 0 {
		t.Fatal("missing user agent")
	}
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient()
	fmt.Println(client.UserAgent)
	// Output:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient()
	}
}

// TestClient_GetAddress tests the GetAddress()
func TestClient_GetAddress(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries
	client, err := NewClient()
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
	client, err := NewClient()
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
	client, err := NewClient()
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

// ExampleClient_GetAddress example using GetAddress()
func ExampleClient_GetAddress() {
	client, _ := NewClient()
	resp, _ := client.GetAddress("16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA") //mrz@moneybutton.com
	fmt.Println(resp.Address)
	// Output:16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA
}
