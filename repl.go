package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) == 0 {
			continue
		}
		commandName := cleanText[0]
		args := []string{}
		if len(cleanText) > 1 {
			args = cleanText[1:]
		}
		commands := getCommands()

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area_name}",
			description: "List all pokemons in location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	cleanText := strings.Split(lowerText, " ")
	//cleanText := strings.Fields(lowerText, " ")
	//can use function above
	var result []string

	for i := range cleanText {
		if cleanText[i] == "" {
			continue
		}
		result = append(result, cleanText[i])
	}

	return result
}
