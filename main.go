package main

import (
	"github.com/xaaaaaanny/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	pokemonsCatched     map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{pokeapiClient: pokeapi.NewClient(time.Hour),
		pokemonsCatched: make(map[string]pokeapi.Pokemon)}
	startRepl(&cfg)
}
