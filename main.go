package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	for {
		fmt.Printf("Pokedex >")
		if scanner.Scan() {
			text := scanner.Text()
			new_text := cleanInput(text)

			if len(new_text) == 0 {
				continue
			}
			cmd := new_text[0]
			command, ok := commands[cmd]
			if !ok {
				fmt.Println("Unknown command:", cmd)
				continue
			}
			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}

}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}
