package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dreamsxin/selling-partner-api-go/auth"
)

type contextKey string

const accessTokenKey contextKey = "amazon-sp-api-access-token"

// WithAccessToken overrides the normal LWA token for one request. Use it for
// an RDT and never reuse the RDT for unrelated requests.
func WithAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, accessTokenKey, token)
}

type Transport struct {
	Base               http.RoundTripper
	TokenSource        auth.TokenSource
	UserAgent          string
	Clock              func() time.Time
	MaxRetries         int
	Backoff            func(attempt int) time.Duration
	RetryNonIdempotent bool
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("amazon sp-api: nil request")
	}
	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}
	clock := t.Clock
	if clock == nil {
		clock = time.Now
	}
	maxRetries := t.MaxRetries
	if maxRetries < 0 {
		maxRetries = 0
	}
	for attempt := 0; ; attempt++ {
		current, err := cloneRequest(req)
		if err != nil {
			return nil, err
		}
		if err := t.decorate(current, clock); err != nil {
			return nil, err
		}
		resp, err := base.RoundTrip(current)
		if err != nil {
			if attempt < maxRetries && retryableMethod(req.Method, t.RetryNonIdempotent) {
				if err := sleepContext(req.Context(), t.delay(attempt)); err != nil {
					return nil, err
				}
				continue
			}
			return nil, err
		}
		if !shouldRetry(resp.StatusCode) || attempt >= maxRetries || !retryableMethod(req.Method, t.RetryNonIdempotent) {
			return resp, nil
		}
		if resp.Body != nil {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
		}
		if err := sleepContext(req.Context(), retryAfter(resp.Header, t.delay(attempt))); err != nil {
			return nil, err
		}
	}
}

func (t *Transport) decorate(req *http.Request, clock func() time.Time) error {
	if t.TokenSource != nil {
		token, ok := req.Context().Value(accessTokenKey).(string)
		var err error
		if !ok || token == "" {
			token, err = t.TokenSource.AccessToken(req.Context())
			if err != nil {
				return err
			}
		}
		req.Header.Set("x-amz-access-token", token)
	}
	if req.Header.Get("x-amz-date") == "" {
		req.Header.Set("x-amz-date", clock().UTC().Format("20060102T150405Z"))
	}
	if t.UserAgent != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", t.UserAgent)
	}
	return nil
}

func (t *Transport) delay(attempt int) time.Duration {
	if t.Backoff != nil {
		return t.Backoff(attempt)
	}
	base := 250 * time.Millisecond
	for i := 0; i < attempt; i++ {
		base *= 2
	}
	if base > 8*time.Second {
		base = 8 * time.Second
	}
	return base + time.Duration(rand.Int63n(int64(100*time.Millisecond)))
}

func cloneRequest(req *http.Request) (*http.Request, error) {
	clone := req.Clone(req.Context())
	if req.Body == nil {
		return clone, nil
	}
	if req.GetBody == nil {
		return nil, fmt.Errorf("amazon sp-api: request body cannot be replayed; set GetBody or disable retries")
	}
	body, err := req.GetBody()
	if err != nil {
		return nil, err
	}
	clone.Body = body
	return clone, nil
}

func retryableMethod(method string, allowNonIdempotent bool) bool {
	if allowNonIdempotent {
		return true
	}
	switch strings.ToUpper(method) {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return true
	default:
		return false
	}
}

func shouldRetry(status int) bool { return status == http.StatusTooManyRequests || status >= 500 }

func retryAfter(header http.Header, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(header.Get("Retry-After"))
	if value == "" {
		return fallback
	}
	if seconds, err := strconv.Atoi(value); err == nil && seconds >= 0 {
		return time.Duration(seconds) * time.Second
	}
	if when, err := http.ParseTime(value); err == nil {
		if delay := time.Until(when); delay > 0 {
			return delay
		}
	}
	return fallback
}

func sleepContext(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return nil
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
