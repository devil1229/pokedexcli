package main

import (
	"encoding/json"
	"fmt"

	"github.com/devil1229/pokedexcli/internal/pokeapi"
)

//callback func for map command
func commandMap(currentConfig *config) error {
	//get the list of location areas
	var response pokeapi.LocationAreaResponse
	if currentConfig.nextLocationAreaURL != nil {
		data, ok := currentConfig.pokecache.Get(*currentConfig.nextLocationAreaURL)
		if ok {
			locationAreaResponse := pokeapi.LocationAreaResponse{}
			// Unmarshal the JSON data into the LocationAreaResponse struct
			err := json.Unmarshal(data, &locationAreaResponse)
			if err != nil {
				return err
			}
            response = locationAreaResponse
		} else {
			resp, err := currentConfig.pokeapiClient.ListLocationAreas(currentConfig.nextLocationAreaURL)
			if err != nil {
				return err
			}
			jsonData, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			currentConfig.pokecache.Add(*currentConfig.nextLocationAreaURL, jsonData)
			response = resp
		}
	} else {

		resp, err := currentConfig.pokeapiClient.ListLocationAreas(currentConfig.nextLocationAreaURL)
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		fmt.Println("error due to ")
		fmt.Print(currentConfig.pokecache)
		currentConfig.pokecache.Add("", jsonData)
		fmt.Println("no error due to ")
		response = resp
	}
	
	fmt.Println("Location Areas: ")
	
	//update the nextpage and previous page url
	currentConfig.nextLocationAreaURL = response.Next
	currentConfig.previousLocationAreaURL = response.Previous
	for _, result := range response.Results {
		fmt.Println(" - " , result.Name)
	}
	return nil
}

//callback func for mapb command
func commandMapB(currentConfig *config) error {
	if currentConfig.previousLocationAreaURL == nil {
		fmt.Println("Wrong Input")
		fmt.Println("Please use \"map\" command first and then \"mapb\"")
		return nil
	}
	var response pokeapi.LocationAreaResponse
	data, ok := currentConfig.pokecache.Get(*currentConfig.previousLocationAreaURL)
	if ok {
		locationAreaResponse := pokeapi.LocationAreaResponse{}
		// Unmarshal the JSON data into the LocationAreaResponse struct
		err := json.Unmarshal(data, &locationAreaResponse)
		if err != nil {
			return err
		}
		response = locationAreaResponse
	} else {
		resp, err := currentConfig.pokeapiClient.ListLocationAreas(currentConfig.previousLocationAreaURL)
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		currentConfig.pokecache.Add(*currentConfig.previousLocationAreaURL, jsonData)
		response = resp
	}
	
	fmt.Println("Location Areas: ")
	currentConfig.nextLocationAreaURL = response.Next
	currentConfig.previousLocationAreaURL = response.Previous
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}
