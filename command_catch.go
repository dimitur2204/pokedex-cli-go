package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_BASE_EXP = 600
)

func commandCatch(cfg *config, args []string) error {

	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(args[0], &cfg.pokeCache)
	if err != nil {
		return err
	}

	randSeed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSeed)
	catchChance := (MAX_BASE_EXP / pokemonResp.BaseExperience) * 10
	if catchChance > r.Intn(100) {
		fmt.Println("You caught a " + pokemonResp.Name + "!")
		cfg.pokedex[pokemonResp.Name] = pokemonResp
		fmt.Println("You have " + fmt.Sprint(len(cfg.pokedex)) + " pokemon in your pokedex")
	} else {
		fmt.Println("The " + pokemonResp.Name + " got away!")
	}
	return nil
}
