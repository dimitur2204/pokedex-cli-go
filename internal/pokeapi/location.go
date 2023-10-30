package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dimitur2204/pokedex-cli-go/internal/pokeapi/internal/pokecache"
)

// ListLocations -
func (c *Client) ExploreLocation(name string, cache *pokecache.Cache) (RespLocation, error) {
	url := baseURL + "/location-area/" + name
	fmt.Println(url)
	cached, ok := cache.Get(url)

	locationsResp := RespLocation{}
	if ok {
		fmt.Println("Cachehit")
		err := json.Unmarshal(cached, &locationsResp)
		if err != nil {
			return RespLocation{}, err
		}
		return locationsResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocation{}, err
	}

	cache.Set(url, dat)
	return locationsResp, nil
}
