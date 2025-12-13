package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("You have not caught any pokemon yet")
	}
	fmt.Println("Your Pokedex: ")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println(" - ", pokemon.Name)
	}
	return nil
}
