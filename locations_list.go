package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResponseLocations, error) {
	// Default URL
	url := baseURL + "/location-area"
	// Update to unique URL if passed in
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocations{}, err
	}

	locationsResp := ResponseLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return ResponseLocations{}, err
	}

	return locationsResp, nil
}
