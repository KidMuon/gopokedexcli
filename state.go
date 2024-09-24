package main

import (
	"github.com/KidMuon/gopokedexcli/internal/pokecache"
	"time"
)

type replState struct {
	PokeLocationNextUrl  string
	PokeLocationPrevUrl  string
	Cache                pokecache.Cache
	PokeEncounterBaseUrl string
}

func initialState() *replState {
	state := replState{
		PokeLocationNextUrl:  "https://pokeapi.co/api/v2/location-area/",
		Cache:                *pokecache.NewCache(5 * time.Minute),
		PokeEncounterBaseUrl: "https://pokeapi.co/api/v2/location-area/",
	}
	return &state
}
