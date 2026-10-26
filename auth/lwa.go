package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const DefaultLWAEndpoint = "https://api.amazon.com/auth/o2/token"

// Credentials contains the LWA application credentials for one application.
// RefreshToken is tenant-specific and must not be shared between sellers.
type Credentials struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

// TokenSource supplies an access token for an SP-API request.
type TokenSource interface {
	AccessToken(context.Context) (string, error)
}

type LWAProvider struct {
	HTTP        *http.Client
	Endpoint    string
	Credentials Credentials
	Cache       TokenCache
	Clock       func() time.Time
	RefreshSkew time.Duration
	cacheKey    string
	mu          sync.Mutex
}

func NewLWAProvider(credentials Credentials) *LWAProvider {
	return &LWAProvider{
		HTTP:        http.DefaultClient,
		Endpoint:    DefaultLWAEndpoint,
		Credentials: credentials,
		Cache:       NewMemoryTokenCache(),
		Clock:       time.Now,
		RefreshSkew: time.Minute,
		cacheKey:    credentials.ClientID + ":refresh:" + credentials.RefreshToken,
	}
}

func (p *LWAProvider) AccessToken(ctx context.Context) (string, error) {
	if p.Credentials.RefreshToken == "" {
		return "", errors.New("amazon sp-api: refresh token is required")
	}
	return p.token(ctx, "refresh_token", "", p.Credentials.RefreshToken)
}

// GrantlessToken obtains an application token for a grantless scope.
func (p *LWAProvider) GrantlessToken(ctx context.Context, scope string) (string, error) {
	if strings.TrimSpace(scope) == "" {
		return "", errors.New("amazon sp-api: grantless scope is required")
	}
	return p.token(ctx, "client_credentials", scope, "")
}

func (p *LWAProvider) token(ctx context.Context, grantType, scope, refreshToken string) (string, error) {
	if p.HTTP == nil {
		p.HTTP = http.DefaultClient
	}
	if p.Endpoint == "" {
		p.Endpoint = DefaultLWAEndpoint
	}
	if p.Clock == nil {
		p.Clock = time.Now
	}
	if p.RefreshSkew <= 0 {
		p.RefreshSkew = time.Minute
	}
	key := p.cacheKey
	if grantType == "client_credentials" {
		key = p.Credentials.ClientID + ":grantless:" + scope
	}
	if p.Cache != nil {
		if cached, ok := p.cachedToken(key); ok {
			return cached.AccessToken, nil
		}
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if p.Cache != nil {
		if cached, ok := p.cachedToken(key); ok {
			return cached.AccessToken, nil
		}
	}

	form := url.Values{
		"grant_type":    {grantType},
		"client_id":     {p.Credentials.ClientID},
		"client_secret": {p.Credentials.ClientSecret},
	}
	if scope != "" {
		form.Set("scope", scope)
	}
	if refreshToken != "" {
		form.Set("refresh_token", refreshToken)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.Endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	resp, err := p.HTTP.Do(req)
	if err != nil {
		return "", fmt.Errorf("amazon sp-api: LWA request: %w", err)
	}
	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return "", fmt.Errorf("amazon sp-api: read LWA response: %w", readErr)
	}
	if resp.StatusCode/100 != 2 {
		return "", fmt.Errorf("amazon sp-api: LWA token exchange failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	var payload struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return "", fmt.Errorf("amazon sp-api: decode LWA response: %w", err)
	}
	if payload.AccessToken == "" {
		return "", errors.New("amazon sp-api: LWA response has no access_token")
	}
	expiresIn := time.Duration(payload.ExpiresIn) * time.Second
	if expiresIn <= p.RefreshSkew {
		expiresIn = time.Second
	} else {
		expiresIn -= p.RefreshSkew
	}
	token := Token{AccessToken: payload.AccessToken, ExpiresAt: p.Clock().Add(expiresIn)}
	if p.Cache != nil {
		p.Cache.Set(key, token)
	}
	return token.AccessToken, nil
}

func (p *LWAProvider) cachedToken(key string) (Token, bool) {
	if p.Cache == nil {
		return Token{}, false
	}
	token, ok := p.Cache.Get(key)
	if !ok || token.AccessToken == "" {
		return Token{}, false
	}
	clock := p.Clock
	if clock == nil {
		clock = time.Now
	}
	if !clock().Before(token.ExpiresAt) {
		return Token{}, false
	}
	return token, true
}
