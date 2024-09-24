package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(state *replState, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Incorrect number of arguments to explore.\nexpected 1 got %d", len(args))
	}
	areaName := args[0]
	fullUrl := state.PokeEncounterBaseUrl + areaName

	var data []byte
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

		state.Cache.Add(fullUrl, data)
	}

	var pokemonInArea PokeAPIPokemonInAreaResponse
	if err := json.Unmarshal(data, &pokemonInArea); err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", areaName)
	for _, encounter := range pokemonInArea.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

type PokeAPIPokemonInAreaResponse struct {
	Encounters []PokeAPIEncounters `json:"pokemon_encounters"`
}

type PokeAPIEncounters struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}
