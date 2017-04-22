package bicimad_fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/ajnavarro/bicimad-fetcher/model"
	"net/http"
	"time"
)

const (
	getStationsURL = "https://rbdata.emtmadrid.es:8443/BiciMad/get_stations/%s/%s"

	timeout = 30 * time.Second
)

type Client struct {
	id  string
	key string

	c *http.Client
}

func NewClient(id string, key string) *Client {
	return &Client{id, key,
		&http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) Stations() (*model.Stations, error) {
	resp, err := c.c.Get(fmt.Sprintf(getStationsURL, c.id, c.key))
	if err != nil {
		return nil, err
	}

	result := &model.Result{}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}

	stations := &model.Stations{}
	if err := json.Unmarshal([]byte(result.Data), stations); err != nil {
		return nil, err
	}

	return stations, nil
}
