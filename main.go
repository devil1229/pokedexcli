package main

import "github.com/devil1229/pokedexcli/internal/pokeapi"


//config needed across the application will be defined here
type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL          *string
	previousLocationAreaURL      *string
}

func main() {
    cfg := config {
		pokeapiClient: pokeapi.NewClient(),
	}
	//start the repl
	startrepl(&cfg)
}
