package main

import (
	"fmt"
)

//callback func for map command
func commandMap(currentConfig *config) error {
	//get the list of location areas
	response, err := currentConfig.pokeapiClient.ListLocationAreas(currentConfig.nextLocationAreaURL)
	if err != nil {
		return err
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
	response, err := currentConfig.pokeapiClient.ListLocationAreas(currentConfig.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")
	currentConfig.nextLocationAreaURL = response.Next
	currentConfig.previousLocationAreaURL = response.Previous
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}
