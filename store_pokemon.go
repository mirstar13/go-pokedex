package main

func storePokemon(cfg *config, pokemonName string) error {
	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	cfg.caughtPokemon[pokemonName] = pokemonInfo
	return nil
}
