package main

import "fmt"

func commandPokedex(currentConfig *config, name string) error {
	
	fmt.Println("Your Pokedex: ")
	for key :=  range currentConfig.pokemonData {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}