package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(currentConfig *config, name string) error {

	if name == "" {
		return fmt.Errorf("please provide the Pokemon name as ARG")
	}
	response, err := currentConfig.pokeapiClient.GetPokemonDetails(name, currentConfig.pokecache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	//baseExperience := response.BaseExperience
	randomNumber := rand.Intn(10)
	if randomNumber != 9 {
		fmt.Printf("%v escaped!\n", name)
	} else {
		fmt.Printf("%v was caught!\n", name)
        currentConfig.pokemonData[name] = response
	}
	return nil
}