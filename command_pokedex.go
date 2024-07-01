package main

import "fmt"

func commandPokedex(cfg *config, s string) error {
	fmt.Println("Your Pokedex:")
	for pokemonNames := range cfg.caughtPokemon {
		fmt.Printf(" -%s\n", pokemonNames)
	}
	return nil
}
