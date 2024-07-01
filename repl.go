package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MrAinslay/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, s string) error
}

func cleanInput(s string) string {
	output := strings.TrimSpace(s)
	output = strings.ToLower(output)
	return output
}

func printPrompt() {
	fmt.Print("Pokedex > ")
}

func startRpl(cfg *config) {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		splitText := strings.Split(text, " ")
		if command, exists := commands[splitText[0]]; exists {
			if len(splitText) > 1 {
				err := command.callback(cfg, splitText[1])
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err := command.callback(cfg, splitText[0])
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("Unknown Command")
		}
		printPrompt()
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location to see a list of Pokemon in the area(usage: explore <location_name/location_id>)",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon(usage: catch <pokemon_name>)",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the name, height, weight, stats and type(s) of a pokemon(usage: inspect <pokemon_name>)",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
