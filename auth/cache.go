package auth

import (
	"sync"
	"time"
)

// Token is an access token returned by Login With Amazon (LWA).
type Token struct {
	AccessToken string
	ExpiresAt   time.Time
}

// TokenCache stores short-lived access tokens. Implementations should keep
// refresh tokens outside this cache and in a secret manager or encrypted store.
type TokenCache interface {
	Get(key string) (Token, bool)
	Set(key string, token Token)
}

// MemoryTokenCache is a process-local cache suitable for one application
// instance. Use a shared cache when multiple replicas need to reuse tokens.
type MemoryTokenCache struct {
	mu     sync.RWMutex
	tokens map[string]Token
}

func NewMemoryTokenCache() *MemoryTokenCache {
	return &MemoryTokenCache{tokens: make(map[string]Token)}
}

func (c *MemoryTokenCache) Get(key string) (Token, bool) {
	c.mu.RLock()
	token, ok := c.tokens[key]
	c.mu.RUnlock()
	if !ok || token.AccessToken == "" {
		return Token{}, false
	}
	return token, true
}

func (c *MemoryTokenCache) Set(key string, token Token) {
	c.mu.Lock()
	if c.tokens == nil {
		c.tokens = make(map[string]Token)
	}
	c.tokens[key] = token
	c.mu.Unlock()
}
