package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/devil1229/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocationAreas(pageURL *string, cache *pokecache.Cache) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	var data []byte

	if pageURL != nil {
		fullURL = *pageURL
	}
	data1, ok := cache.Get(fullURL)
	if ok { 
       data = data1
	} else {
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
		data2, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
        cache.Add(fullURL, data2)
		data = data2
	}
		
	locationAreaResponse := LocationAreaResponse{}
	// Unmarshal the JSON data into the LocationAreaResponse struct
	//fmt.Println("Data From cache: " , string(data))
	err := json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse , nil
}