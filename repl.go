package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	pokeapiClient   Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		commandName := reader.Text()

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type CliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 locations in the Pokemon world. Subsequent calls return the next 20 locations.",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 previous locations. Subsequent calls return the next previous 20 locations.",
			callback:    mapbCommand,
		},
	}
}
