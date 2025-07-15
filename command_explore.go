package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area, e.g., 'explore pastoria-city-area'")
	}

	url := baseURL + args[0] //Add URL here

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data pokemonAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	fmt.Printf("Exploring %s...\n", args[0])
	for _, encounter := range data.PokemonEncounter {
		fmt.Println("- ", encounter.Pokemon.Name)
	}

	return nil
}
