package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands:")

	commands := getCommands()
	for k := range commands {
		fmt.Printf("~%v: %v\n", commands[k].name, commands[k].description)
	}
	return nil
}
