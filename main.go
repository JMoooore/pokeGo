package main

import (
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

func main() {
	pokeClient := NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
