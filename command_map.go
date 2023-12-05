package main

import (
	"fmt"

	"github.com/devil1229/pokedexcli/internal/pokeapi"
)

func commandMap(currentConfig *config) error {
	// fmt.Print(currentConfig)
	response, err := pokeapi(currentConfig.next)

	if err != nil {
		return err
	}
	currentConfig.next = response.Next
	currentConfig.previous = response.Previous
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(currentConfig *config) error {
	// fmt.Print(currentConfig)
	if currentConfig.previous == nil {
		fmt.Println("Wrong Input")
		fmt.Println("Please use \"map\" command first and then \"mapb\"")
		return nil
	}
	response, err := pokeapi(*currentConfig.previous)

	if err != nil {
		return err
	}
	currentConfig.next = response.Next
	currentConfig.previous = response.Previous
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}
