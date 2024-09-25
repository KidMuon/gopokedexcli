package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startPrompt() {
	validCommands := getCommands()
	state := initialState()
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string
	for {
		fmt.Println()
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

		if len(userCommand) > 1 {
			call(validCommands[userCommand[0]].callback, state, userCommand[1:])
		} else {
			call(validCommands[userCommand[0]].callback, state, []string{})
		}

	}
}

func cleanAndParseCommand(input string) []string {
	lowercase := strings.ToLower(input)
	words := strings.Fields(lowercase)
	return words
}

func call(callback func(*replState, []string) error, state *replState, args []string) {
	err := callback(state, args)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*replState, []string) error
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
		"map": {
			name:        "map",
			description: "Display the next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows the pokemon that can be encountered in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the named pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show details of the named pokemon if you've caught them",
			callback:    commandInspect,
		},
	}
	return commands
}
