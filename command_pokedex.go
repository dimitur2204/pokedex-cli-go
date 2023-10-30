package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	for _, v := range cfg.pokedex {
		fmt.Println("Name: " + v.Name)
	}
	return nil
}
