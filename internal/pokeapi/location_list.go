package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/dimitur2204/pokedex-cli-go/internal/pokeapi/internal/pokecache"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	mux := &sync.Mutex{}
	cached, ok := cache.Get(url, mux)

	locationsResp := RespShallowLocations{}
	if ok {
		fmt.Println("Cachehit")
		err := json.Unmarshal(*cached, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	cache.Set(url, &dat, mux)
	return locationsResp, nil
}
