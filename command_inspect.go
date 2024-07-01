package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, s string) error {
	pokemon, ok := cfg.caughtPokemon[s]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("  -%s\n", types.Type.Name)
	}
	return nil
}
