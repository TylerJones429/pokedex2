package main

import (
	"time"

	"github.com/TylerJones429/pokedex2/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
