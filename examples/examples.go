/*
Package main is an example package using go-polynym
*/
package main

import (
	"log"

	"github.com/mrz1836/go-polynym"
)

func main() {

	// Start a new client
	client, err := polynym.NewClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Resolve a handle
	var resp *polynym.GetAddressResponse
	resp, err = client.GetAddress("mrz@moneybutton.com")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Success
	log.Println("address:", resp.Address)
}
