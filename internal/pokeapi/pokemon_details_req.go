package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/devil1229/pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemonDetails(name string, cache *pokecache.Cache) (Pokemon, error) {
	endpoint := "pokemon/" + name
	fullURL := baseURL + endpoint
	var data []byte
    
	//checking the cache hit
	data1, ok := cache.Get(fullURL)
	if ok {
		data = data1
	} else {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return Pokemon{}, err
		}

		//calling the http request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 399 {
			return Pokemon{} , fmt.Errorf("bad Status Code: %v", resp.StatusCode)
		}

		//reading the data from the response body -- basically converting it into slice of bytes 
		//which unmarshal func can take as a input
		data2, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
		//catching the data
		cache.Add(fullURL, data2)
		data = data2
	}

	pokemon := Pokemon{}
	// Unmarshal the JSON data into the LocationAreaResponse struct
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon , nil	
}