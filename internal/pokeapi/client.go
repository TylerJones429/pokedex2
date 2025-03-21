package pokeapi

import (
	"net/http"
	"pokedex2/internal/pokecache"
	"time"
)

// Client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// newClient
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
