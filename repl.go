package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) == 0 {
			continue
		}
		commandName := cleanText[0]
		commands := getCommands()

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
