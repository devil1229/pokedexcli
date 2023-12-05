package main

import "fmt"

func commandHelp(currentConfig *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	allCommands := getALlCliCommands()

	for key, val := range allCommands {
		fmt.Printf("%v: %v", key, val.description)
		fmt.Println()
	}
	fmt.Println()
	return nil
}
