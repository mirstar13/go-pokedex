package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(locationName string) (Encounters, error) {
	url := baseURL + "/location-area/"
	url += locationName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Encounters{}, err
	}

	if val, ok := c.cache.Get(url); ok {
		encountersResp := Encounters{}
		err := json.Unmarshal(val, &encountersResp)
		if err != nil {
			return Encounters{}, err
		}

		return encountersResp, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Encounters{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Encounters{}, err
	}

	encountersResp := Encounters{}
	err = json.Unmarshal(dat, &encountersResp)
	if err != nil {
		return Encounters{}, err
	}

	c.cache.Add(url, dat)
	return encountersResp, nil
}
