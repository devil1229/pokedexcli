package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreaResponse, error) {
	endpoint := "/Location"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
		return LocationAreaResponse{}, err
	}
    
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResponse{} , fmt.Errorf("bad Status Code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse , nil
}