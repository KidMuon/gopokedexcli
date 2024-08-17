package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startPrompt() {
	validCommands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string
	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		userInput = scanner.Text()

		userCommand := cleanAndParseCommand(userInput)
		if len(userCommand) == 0 {
			continue
		}

		if _, ok := validCommands[userCommand[0]]; !ok {
			fmt.Println("Invalid input. Type \"help\" for help.")
			continue
		}

		call(validCommands[userInput].callback)

	}
}

func cleanAndParseCommand(input string) []string {
	lowercase := strings.ToLower(input)
	words := strings.Fields(lowercase)
	return words
}

func call(callback func() error) {
	err := callback()
	if err != nil {
		log.Fatal(err)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
}
