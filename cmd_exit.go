package main

import (
	"fmt"
	"os"
)

func exitCommand(cfg *config) error {
	fmt.Println("Closing the Pokedex!")
	os.Exit(0)
	return nil
}
