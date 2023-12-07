package main

import "fmt"

func commandInspect(currentConfig *config, name string) error {
	if name == "" {
		return fmt.Errorf("please provide the Pokemon name as ARG")
	}

	value, exists := currentConfig.pokemonData[name]
	if exists {
		fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", value.Name, value.Height, value.Weight)
		for _, val := range value.Stats {
			fmt.Printf("  - %v: %v\n", val.Stat.Name, val.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, val := range value.Types {
			fmt.Printf("  - %v\n", val.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}