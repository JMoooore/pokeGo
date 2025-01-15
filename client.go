package main

import (
	"net/http"
	"time"
)

type Client struct {
	cache      Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
