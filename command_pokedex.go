package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokemonsCatched) == 0 {
		return fmt.Errorf("your Pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.pokemonsCatched {
		fmt.Printf("	- %v\n", v.Name)
	}
	return nil
}
