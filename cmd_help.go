package main

import "fmt"

func helpCommand(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex CLI!")
	fmt.Println("Usage Below:")
	fmt.Println()
	fmt.Println("Command: Description")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
