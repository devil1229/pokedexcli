package main

import (
	"fmt"
)

func commandExplore(currentConfig *config, location string) error {	
	
	if location == "" {
		return fmt.Errorf("please provide the location ARG")
	}
	response, err := currentConfig.pokeapiClient.GetPokemonEncouters(location, currentConfig.pokecache)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", location)
    fmt.Println("Found Pokemon: ")

	for _, result := range response.PokemonEncounters {
		fmt.Println(" - " , result.Pokemon.Name)
	}
	return nil
}