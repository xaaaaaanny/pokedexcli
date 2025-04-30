package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	commands := getCommands()
	for k := range commands {
		fmt.Printf("%v: %v\n", commands[k].name, commands[k].description)
	}
	return nil
}
