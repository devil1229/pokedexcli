package main

import "fmt"

func commandHelp(currentConfig *config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	allCommands := getALlCliCommands()

	for _, val := range allCommands {
		fmt.Printf("%v:  %v", val.name, val.description)
		fmt.Println()
	}
	fmt.Println()
	return nil
}
