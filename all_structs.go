package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
}

type config struct {
	nextURL     *string
	previousURL *string
}

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type pokemonAreaResponse struct {
	PokemonEncounter []PokemonEncounter `json:"pokemon_encounters"`
}
