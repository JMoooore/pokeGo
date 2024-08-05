package main

import (
	"errors"
	"fmt"
)

func mapCommand(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = &resp.Next
	cfg.prevLocationURL = &resp.Previous

	for _, locResult := range resp.Results {
		fmt.Println(locResult.Name)
	}
	return nil
}

func mapbCommand(cfg *config) error {
	if cfg.prevLocationURL == nil || *cfg.prevLocationURL == "" {
		return errors.New("You are on the first page!")
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = &resp.Next
	cfg.prevLocationURL = &resp.Previous

	for _, locResult := range resp.Results {
		fmt.Println(locResult.Name)
	}
	return nil
}
