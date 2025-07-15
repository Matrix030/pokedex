package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	cfg := &config{}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if commandName == "explore" {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
		} else if exists {
			args = []string{}
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func(cfg *config, args []string) error {
				return commandHelp()
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func(cfg *config, args []string) error {
				return commandExit()
			},
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the Previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the Pokemon in the location",
			callback:    commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Used to catch a pokemon to add them to the user's pokedex",
			callback: commandCatch,
		}
	}
}
