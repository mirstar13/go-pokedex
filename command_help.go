package main

import "fmt"

func commandHelp(cfg *config, s string) error {
	fmt.Println("\nWelcome to the Pokedex\nUsage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
