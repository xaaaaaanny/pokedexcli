package main

import (
	"github.com/xaaaaaanny/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{pokeapiClient: pokeapi.NewClient(time.Hour)}
	startRepl(&cfg)
}
