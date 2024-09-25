package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
)

type CatchPokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func commandCatch(state *replState, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("too many arguments passed")
	} else if len(args) < 1 {
		return fmt.Errorf("catch must be called with a Pokemon name as argument")
	}

	pokemonName := args[0]
	fullUrl := state.PokePokemonUrl + pokemonName
	data, ok := state.Cache.Get(fullUrl)
	if !ok {
		resp, err := http.Get(fullUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	var pokemon CatchPokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return err
	}

	fmt.Printf("Throwing Pokeball at %s...\n", pokemon.Name)
	if catch(pokemon.BaseExperience) {
		fmt.Printf("Caught %s!", pokemon.Name)
	} else {
		fmt.Printf("%s got away.", pokemon.Name)
	}

	return nil
}

func catch(baseExperience int) bool {
	var x float64 = float64(baseExperience)
	//Asssume the catch chance follows a logistic function
	// L / ( 1 + e ^ ( -k * (x - x0)))
	//Maximum catch chance is 99.5% represented by 9950 / 10000
	//Minimum catch chance is 0.5% represented by 50 / 10000
	//Other constants determined by eyeballing
	//I used caterpie as my low-end base_experience
	//and arceus as my high-end base_experience
	L := 9900.0
	x0 := 182.0
	y := 50
	k := 0.025

	t := -k * (x - x0)
	denom := 1.0 + math.Exp(t)
	catchValue := int(L/denom) + y

	executionValue := rand.Intn(10000)
	fmt.Println(catchValue, executionValue)

	return executionValue > catchValue
}
