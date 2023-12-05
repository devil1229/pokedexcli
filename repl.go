package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var currentConfig = config{}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startrepl() {

	// Create a scanner for standard input (os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt the user for input
		fmt.Print("Pokedex > ")

		// Scan for the next token (line) from the input
		scanner.Scan()
		// Retrieve the scanned text
		input := scanner.Text()
        // Split the input into commands based on spaces by cleaninput function
		commands := cleanInput(input)
        
		//check if user has entered a command or not, if not then continue taking the input again
		if len(commands) == 0 {
			continue
		}

		availableCommands := getALlCliCommands()
		commandName := commands[0]
        
        command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Wrong Input!")
			fmt.Println("please type \"help\" to see working commands")
			continue
		}

		err := command.callback(&currentConfig)

		if err != nil {
			fmt.Printf("Something Went Wrong : %v\n", err)
			continue
		}	
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}


func getALlCliCommands() map[string]cliCommand {

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