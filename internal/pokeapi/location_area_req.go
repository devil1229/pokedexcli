package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}
	//creating a new http request 
	req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
		return LocationAreaResponse{}, err
	}
    
	//calling the http request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResponse{} , fmt.Errorf("bad Status Code: %v", resp.StatusCode)
	}

	//reading the data from the response body -- basically converting it into slice of bytes 
	//which unmarshal func can take as a input
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	// Unmarshal the JSON data into the LocationAreaResponse struct
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse , nil
}