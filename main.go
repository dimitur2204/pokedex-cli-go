package main

import (
	"time"

	"github.com/dimitur2204/pokedex-cli-go/internal/pokeapi/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
