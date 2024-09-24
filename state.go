package main

import (
	"github.com/KidMuon/gopokedexcli/internal/pokecache"
	"time"
)

type replState struct {
	PokeLocationNextUrl string
	PokeLocationPrevUrl string
	Cache               pokecache.Cache
}

func initialState() *replState {
	state := replState{
		PokeLocationNextUrl: "https://pokeapi.co/api/v2/location/",
		Cache:               *pokecache.NewCache(5 * time.Minute),
	}
	return &state
}
