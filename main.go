package main

import (
	"time"

	"github.com/devil1229/pokedexcli/internal/pokeapi"
	"github.com/devil1229/pokedexcli/internal/pokecache"
)

const interval = 5 * time.Second

//config needed across the application will be defined here
type config struct {
	pokeapiClient pokeapi.Client
	pokemonData map[string]pokeapi.Pokemon
	pokecache *pokecache.Cache
	nextLocationAreaURL          *string
	previousLocationAreaURL      *string
}

func main() {
    cfg := config {
		pokeapiClient: pokeapi.NewClient(),
		pokecache: pokecache.NewCache(interval),
		pokemonData: map[string]pokeapi.Pokemon{},
	}
	//start the repl
	startrepl(&cfg)
}
