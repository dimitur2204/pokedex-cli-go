package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {

	locationResp, err := cfg.pokeapiClient.ExploreLocation(args[0], &cfg.pokeCache)
	if err != nil {
		return err
	}

	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}
