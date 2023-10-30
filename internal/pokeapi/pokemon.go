package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dimitur2204/pokedex-cli-go/internal/pokeapi/internal/pokecache"
)

func (c *Client) CatchPokemon(name string, cache *pokecache.Cache) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name
	fmt.Println(url)
	cached, ok := cache.Get(url)

	pokemonResp := RespPokemon{}
	if ok {
		fmt.Println("Cachehit")
		err := json.Unmarshal(cached, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	cache.Set(url, dat)
	return pokemonResp, nil
}
