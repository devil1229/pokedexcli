package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/devil1229/pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemonEncouters(location string, cache *pokecache.Cache) (PokemonEncounters, error){

	endpoint := "location-area/" + location
	fullURL := baseURL + endpoint
	var data []byte

	data1, ok := cache.Get(fullURL)
	if ok {
		data = data1
	} else {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return PokemonEncounters{}, err
		}

		//calling the http request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonEncounters{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 399 {
			return PokemonEncounters{} , fmt.Errorf("bad Status Code: %v", resp.StatusCode)
		}

		//reading the data from the response body -- basically converting it into slice of bytes 
		//which unmarshal func can take as a input
		data2, err := io.ReadAll(resp.Body)
		if err != nil {
			return PokemonEncounters{}, err
		}
		cache.Add(fullURL, data2)
		data = data2
	}

	pokemonEncounters := PokemonEncounters{}
	// Unmarshal the JSON data into the LocationAreaResponse struct
	err := json.Unmarshal(data, &pokemonEncounters)
	if err != nil {
		return PokemonEncounters{}, err
	}

	return pokemonEncounters , nil
}