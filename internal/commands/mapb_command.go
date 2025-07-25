package commands

import (
	"fmt"
	"github.com/D3rise/pokedexcli/pokeapi"
	"log"

	"github.com/D3rise/pokedexcli/internal/context"
)

var mapbCommandMeta = cliCommand{
	Name:        "mapb",
	Args:        []string{},
	Description: "Print locations, each usage decreases offset by 20",
	Callback:    mapbCommand,
}

const mapPreviousOffset context.ContextKey = "mapPreviousOffset"

func mapbCommand(c *context.Context, _ ...string) error {
	api, ok := c.Get(pokeapi.PokeApiContextKey).(*pokeapi.PokeAPI)

	if !ok {
		log.Fatal("pokedex api is not initialized")
	}

	if !c.Has(mapPreviousOffset) {
		c.Set(mapPreviousOffset, 0)
	}

	offset := c.Get(mapPreviousOffset).(int)
	if offset < 0 {
		offset = 0
	}

	result, err := api.GetLocationAreaList(20, offset)
	if err != nil {
		log.Fatal(err)
	}

	c.Set(mapPreviousOffset, offset-20)
	c.Set(mapNextOffset, offset+20)

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}
