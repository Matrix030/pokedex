package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Matrix030/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/location-area/"

var cache = pokecache.NewCache(5 * time.Minute)

func commandMap(cfg *config) error {
	url := baseURL
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}

	if cachedData, found := cache.Get(url); found {
		fmt.Println("Using cached data!")

		var data locationAreaResponse
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return err
		}
		for _, area := range data.Results {
			fmt.Println(area.Name)
		}
		cfg.nextURL = data.Next
		cfg.previousURL = data.Previous
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

	cache.Add(url, body)

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
		fmt.Println("You are on the first page")
		return nil
	}

	if cachedData, found := cache.Get(*cfg.previousURL); found {
		var data locationAreaResponse
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return err
		}

		for _, area := range data.Results {
			fmt.Println(area.Name)
		}
		cfg.nextURL = data.Next
		cfg.previousURL = data.Previous
		return nil
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
