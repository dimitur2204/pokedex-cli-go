package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("You have not caught a " + name)
		return nil
	}
	fmt.Println("Name: " + pokemon.Name)
	fmt.Println("Base Experience: " + fmt.Sprint(pokemon.BaseExperience))
	fmt.Println("Height: " + fmt.Sprint(pokemon.Height))
	fmt.Println("Weight: " + fmt.Sprint(pokemon.Weight))
	fmt.Println("Abilities:")
	for _, v := range pokemon.Abilities {
		fmt.Println(v.Ability.Name)
	}
	fmt.Println("Moves:")
	for _, v := range pokemon.Moves {
		fmt.Println(v.Move.Name)
	}
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Println(v.Stat.Name + ": " + fmt.Sprint(v.BaseStat))
	}
	return nil
}
