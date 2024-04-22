package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"tryhardplayoffs/internal/mysportsfeeds"
	"tryhardplayoffs/internal/utils"
)

func main() {
	// Define the API endpoint
	apiURL := "https://api.mysportsfeeds.com/v2.1/pull/nhl/players.json"

	// Create query parameters
	params := url.Values{}
	utils.SetQueryParams(mysportsfeeds.QueryParams{Player: "kucherov"}, &params)

	// Add the query parameters to the URL
	apiURL += "?" + params.Encode()

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the required authentication headers
	req.SetBasicAuth("0a17aeb8-8db7-484a-a42a-dfbf6e", "MYSPORTSFEEDS")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status code:", resp.StatusCode)
		return
	}

	// Parse the JSON response
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the data
	fmt.Println("Response data:", data)
}
