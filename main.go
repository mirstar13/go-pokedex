package main

import (
	"time"

	"github.com/MrAinslay/Pokedex/internal/pokeapi"
)

func main() {
	caughtPokemon := make(map[string]pokeapi.Pokemon)
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: caughtPokemon,
	}
	startRpl(cfg)
}
