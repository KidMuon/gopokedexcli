package main

import "fmt"

func commandPokedex(state *replState, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("pokedex takes no arguments")
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range state.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
