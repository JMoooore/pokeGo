package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display commands for utilizing the CLI tool",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCommand,
		},
	}
}

func helpCommand() error {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")

	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}
