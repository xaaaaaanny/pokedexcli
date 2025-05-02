package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
