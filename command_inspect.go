package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemonName := args[0]
	_, ok := cfg.pokemonsCatched[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon is not caught yet")
	}

	fmt.Println("Name: ", cfg.pokemonsCatched[pokemonName].Name)
	fmt.Println("Weight: ", cfg.pokemonsCatched[pokemonName].Weight)
	fmt.Println("Stats:")
	for _, stat := range cfg.pokemonsCatched[pokemonName].Stats {
		fmt.Printf("	-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range cfg.pokemonsCatched[pokemonName].Types {
		fmt.Printf("	-%v\n", v.Type.Name)
	}
	return nil
}
