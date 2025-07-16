package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	pokeapi "github.com/Matrix030/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	caughtPokemon := map[string]pokeapi.Pokemon{}
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if r.Intn(pokemon.BaseExperience) == 0 {
		fmt.Printf("%s was caught\n", pokemon.Name)
		caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
