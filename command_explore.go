package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no location area provided")
	}
	locationArea := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + locationArea)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println("-" + pokemon.Pokemon.Name)
	}

	return nil
}
