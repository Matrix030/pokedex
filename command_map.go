package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2/location-area/"

func commandMap(cfg *config) error {
	url := baseURL
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return err2
	}

	var data locationAreaResponse
	if err3 := json.Unmarshal(body, &data); err != nil {
		return err3
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	cfg.nextURL = data.Next
	cfg.previousURL = data.Previous

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("You are you on the first page")
	}

	resp, err := http.Get(*cfg.previousURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err2 := io.ReadAll(resp.Body)

	if err2 != nil {
		return err2
	}
	var data locationAreaResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	cfg.nextURL = data.Next
	cfg.previousURL = data.Previous

	return nil
}
