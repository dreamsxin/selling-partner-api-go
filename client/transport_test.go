package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type staticToken string

func (s staticToken) AccessToken(context.Context) (string, error) { return string(s), nil }

func TestTransportAddsSPAPIHeadersAndRetries(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if got := r.Header.Get("x-amz-access-token"); got != "token" {
			t.Fatalf("token header = %q", got)
		}
		if r.Header.Get("x-amz-date") == "" || r.Header.Get("User-Agent") != "ERP/1.0" {
			t.Fatalf("missing SP-API headers: %#v", r.Header)
		}
		if attempts == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		_, _ = io.WriteString(w, `{"ok":true}`)
	}))
	defer server.Close()
	httpClient := NewHTTPClient(HTTPClientOptions{TokenSource: staticToken("token"), UserAgent: "ERP/1.0", MaxRetries: 1})
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if attempts != 2 || resp.StatusCode != http.StatusOK {
		t.Fatalf("attempts=%d status=%d", attempts, resp.StatusCode)
	}
}

func TestRetryDoesNotReplayPostByDefault(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	httpClient := NewHTTPClient(HTTPClientOptions{TokenSource: staticToken("token"), MaxRetries: 3})
	req, _ := http.NewRequest(http.MethodPost, server.URL, strings.NewReader("body"))
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if attempts != 1 {
		t.Fatalf("POST attempts=%d", attempts)
	}
}

func TestRetryContextCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if err := sleepContext(ctx, time.Second); err == nil {
		t.Fatal("expected context cancellation")
	}
}
