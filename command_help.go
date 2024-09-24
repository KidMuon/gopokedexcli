package main

import "fmt"

func commandHelp(state *replState) error {
	commands := getCommands()
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, v := range commands {
		_, err := fmt.Printf("%s: %s\n", v.name, v.description)
		if err != nil {
			return err
		}
	}
	return nil
}
