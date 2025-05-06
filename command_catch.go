package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemonName := args[0]
	_, ok := cfg.pokemonsCatched[pokemonName]
	if ok {
		return fmt.Errorf("pokemon already catched")
	}

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	if isCatched(resp.BaseExperience) {
		fmt.Println(pokemonName + " was caught!")
		cfg.pokemonsCatched[pokemonName] = resp
	} else {
		fmt.Println(pokemonName + " escaped")
	}

	return nil
}

func isCatched(baseExperience int) bool {
	threshold := baseExperience / 2
	randNum := rand.Intn(baseExperience)
	if randNum > threshold {
		return true
	}
	return false
}
