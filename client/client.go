package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dreamsxin/selling-partner-api-go/auth"
)

type HTTPClientOptions struct {
	HTTPClient         *http.Client
	TokenSource        auth.TokenSource
	UserAgent          string
	MaxRetries         int
	RetryNonIdempotent bool
}

func NewHTTPClient(options HTTPClientOptions) *http.Client {
	base := http.DefaultTransport
	if options.HTTPClient != nil && options.HTTPClient.Transport != nil {
		base = options.HTTPClient.Transport
	}
	transport := &Transport{
		Base:               base,
		TokenSource:        options.TokenSource,
		UserAgent:          options.UserAgent,
		MaxRetries:         options.MaxRetries,
		RetryNonIdempotent: options.RetryNonIdempotent,
	}
	if options.HTTPClient == nil {
		return &http.Client{Transport: transport, Timeout: 60 * time.Second}
	}
	clone := *options.HTTPClient
	clone.Transport = transport
	return &clone
}

type APIError struct {
	StatusCode int
	RequestID  string
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("amazon sp-api: status=%d request_id=%s body=%s", e.StatusCode, e.RequestID, e.Body)
}

// DoJSON is useful for APIs that have not been generated yet. Generated API
// packages can use the same HTTP client through their Configuration.HTTPClient.
func DoJSON(ctx context.Context, httpClient *http.Client, endpoint, path, method string, query url.Values, body any, out any) error {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	requestURL := strings.TrimRight(endpoint, "/") + "/" + strings.TrimLeft(path, "/")
	parsed, err := url.Parse(requestURL)
	if err != nil {
		return err
	}
	parsed.RawQuery = query.Encode()
	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(encoded)
	}
	req, err := http.NewRequestWithContext(ctx, method, parsed.String(), reader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode/100 != 2 {
		return &APIError{StatusCode: resp.StatusCode, RequestID: resp.Header.Get("x-amzn-RequestId"), Body: strings.TrimSpace(string(responseBody))}
	}
	if out == nil || len(responseBody) == 0 {
		return nil
	}
	return json.Unmarshal(responseBody, out)
}
