package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, s string) error {
	if s == "" {
		return errors.New("you must provide a location name or id")
	}
	encountersResp, err := cfg.pokeapiClient.ListPokemon(s)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", s, "...\nFound Pokemon:")

	for _, enc := range encountersResp.PokemonEncounters {
		fmt.Println("-", enc.Pokemon.Name)
	}
	return err
}
