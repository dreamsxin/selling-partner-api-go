package auth

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLWAProviderCachesRefreshToken(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		body, _ := io.ReadAll(r.Body)
		if !strings.Contains(string(body), "grant_type=refresh_token") {
			t.Fatalf("unexpected form: %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"access_token":"access-1","expires_in":3600}`)
	}))
	defer server.Close()
	provider := NewLWAProvider(Credentials{ClientID: "client", ClientSecret: "secret", RefreshToken: "refresh"})
	provider.Endpoint = server.URL
	provider.Clock = func() time.Time { return time.Unix(100, 0) }
	first, err := provider.AccessToken(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	second, err := provider.AccessToken(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if first != "access-1" || second != first || calls != 1 {
		t.Fatalf("first=%q second=%q calls=%d", first, second, calls)
	}
}

func TestRDTClientUsesLWAAndReturnsRDT(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-amz-access-token") != "lwa" {
			t.Fatalf("token=%q", r.Header.Get("x-amz-access-token"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"restrictedDataToken":"rdt-1"}`)
	}))
	defer server.Close()
	rdtClient := RDTClient{HTTP: server.Client(), Endpoint: server.URL, TokenSource: staticTokenSource("lwa")}
	token, err := rdtClient.CreateRestrictedDataToken(context.Background(), []RestrictedResource{{Method: "GET", Path: "/orders/v0/orders/O/address"}})
	if err != nil {
		t.Fatal(err)
	}
	if token != "rdt-1" {
		t.Fatalf("token=%q", token)
	}
}

type staticTokenSource string

func (s staticTokenSource) AccessToken(context.Context) (string, error) { return string(s), nil }
