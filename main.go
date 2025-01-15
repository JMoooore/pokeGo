package main

import (
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

func main() {
	pokeClient := NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
