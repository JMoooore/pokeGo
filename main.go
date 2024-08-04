package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		commandName := reader.Text()

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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
