package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeAPILocationResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []PokeLocation `json:"results"`
}

type PokeLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func commandMap(state *replState) error {
	if state.PokeLocationNextUrl == "" {
		fmt.Println("No more locations to show")
		return nil
	}

	return getLocationFromAPI(state.PokeLocationNextUrl, state)
}

func commandMapb(state *replState) error {
	if state.PokeLocationPrevUrl == "" {
		fmt.Println("No previous locations to show")
		return nil
	}

	return getLocationFromAPI(state.PokeLocationPrevUrl, state)
}

func getLocationFromAPI(url string, state *replState) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResponse PokeAPILocationResponse
	if err = json.Unmarshal(data, &apiResponse); err != nil {
		return err
	}

	for _, loc := range apiResponse.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	state.PokeLocationNextUrl = apiResponse.Next
	state.PokeLocationPrevUrl = apiResponse.Previous

	return nil
}
