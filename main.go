package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			new_text := cleanInput(text)
			fmt.Println("Your command was:", new_text[0])
		}
	}

}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}
