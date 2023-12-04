package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next     string
	previous *string
}

func getALlCliCommands() map[string]cliCommand {

	currentConfig := config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: nil,
	}

	commandHelp := func() error {
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

	commandExit := func() error {
		os.Exit(0)
		return nil
	}

	commandMap := func() error {
		// fmt.Print(currentConfig)
		response , err := makeApiCall(currentConfig.next)

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


	commandMapB := func() error {
		// fmt.Print(currentConfig)
		if currentConfig.previous == nil {
			fmt.Println("Wrong Input")
			fmt.Println("Please use \"map\" command first and then \"mapb\"")
			return nil
		}
		response , err := makeApiCall(*currentConfig.previous)

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

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of privious 20 location areas. Pokemon's way to go back",
			callback:    commandMapB,
		},
	}
}

func main() {

	//response, _ := makeApiCall("https://pokeapi.co/api/v2/location-area")
	allCommands := getALlCliCommands()

	for {
		fmt.Print("Pokedex > ")
		var input string
		fmt.Scan(&input)

		if _, ok := allCommands[input]; ok {
			err := allCommands[input].callback()

			if err != nil {
				fmt.Println("something wenr wrong : ", err)
				break
			}

		} else {
			fmt.Println("Wrong Input!")
			fmt.Println("please type \"help\" to see working commands")
		}
	}
}
