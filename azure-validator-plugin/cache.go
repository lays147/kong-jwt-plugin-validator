package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type JWKSCache struct {
	jwks *cache.Cache
}

const (
	defaultExpiration = 24 * time.Hour
	purgeTime         = 12 * time.Hour
)

func NewJWKSCache() *JWKSCache {
	Cache := cache.New(defaultExpiration, purgeTime)
	return &JWKSCache{jwks: Cache}
}

func (c *JWKSCache) Set(key string, value interface{}) {
	c.jwks.Set(key, value, cache.DefaultExpiration)
}

func (c *JWKSCache) Get(key string) (interface{}, bool) {
	return c.jwks.Get(key)
}
