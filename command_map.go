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

	data, ok := state.Cache.Get(state.PokeLocationNextUrl)
	if !ok {
		resp, err := http.Get(state.PokeLocationNextUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		state.Cache.Add(state.PokeLocationNextUrl, data)
	} else {
		fmt.Println("Retrieving from cache...")
	}

	var apiResponse PokeAPILocationResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return err
	}

	for _, loc := range apiResponse.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	state.PokeLocationPrevUrl = state.PokeLocationNextUrl
	state.PokeLocationNextUrl = apiResponse.Next

	return nil
}

func commandMapb(state *replState) error {
	if state.PokeLocationPrevUrl == "" {
		fmt.Println("No previous locations to show")
		return nil
	}

	data, ok := state.Cache.Get(state.PokeLocationPrevUrl)
	if !ok {
		resp, err := http.Get(state.PokeLocationPrevUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		state.Cache.Add(state.PokeLocationPrevUrl, data)
	} else {
		fmt.Println("Retrieving from cache...")
	}

	var apiResponse PokeAPILocationResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return err
	}

	for _, loc := range apiResponse.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	state.PokeLocationNextUrl = state.PokeLocationPrevUrl
	state.PokeLocationPrevUrl = apiResponse.Previous

	return nil
}
