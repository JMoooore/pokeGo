package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	prev string
	next string
}

type LocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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

func mapCommand() error {
	resp, err := http.Get(locationUrls.next)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Update previous and next URLs
	locRes := LocationResponse{}
	err = json.Unmarshal(body, &locRes)
	if err != nil {
		fmt.Println(err)
	}
	locationUrls.next = locRes.Next
	locationUrls.prev = locRes.Previous

	for _, locResult := range locRes.Results {
		fmt.Println(locResult.Name)
	}

	return nil
}

func mapbCommand() error {
	resp, err := http.Get(locationUrls.prev)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Update previous and next URLs
	locRes := LocationResponse{}
	err = json.Unmarshal(body, &locRes)
	if err != nil {
		fmt.Println(err)
	}
	locationUrls.next = locRes.Next
	locationUrls.prev = locRes.Previous

	for _, locResult := range locRes.Results {
		fmt.Println(locResult.Name)
	}

	return nil
}
