package main

import (
	"fmt"
)

func commandInspect(state *replState, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect expects a pokemon name to inspect")
	} else if len(args) > 1 {
		return fmt.Errorf("too many arguments passed")
	}

	pokemon, ok := state.CaughtPokemon[args[0]]
	if !ok {
		return fmt.Errorf("you haven't caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		val := stat.BaseStat
		key := stat.Stat.StatName
		fmt.Printf("  -%s: %d\n", key, val)
	}

	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Printf("  - %s\n", poketype.Type.PokemonType)
	}

	return nil
}
