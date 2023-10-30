package main

import (
	"time"

	"github.com/dimitur2204/pokedex-cli-go/internal/pokeapi/internal/pokeapi"
	"github.com/dimitur2204/pokedex-cli-go/internal/pokeapi/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
		pokedex:       make(map[string]pokeapi.RespPokemon),
	}

	startRepl(cfg)
}
