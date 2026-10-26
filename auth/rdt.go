package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const TokensPath = "/tokens/2021-03-01/restrictedDataToken"

type RestrictedResource struct {
	Method       string   `json:"method"`
	Path         string   `json:"path"`
	DataElements []string `json:"dataElements,omitempty"`
}

type RDTClient struct {
	HTTP        *http.Client
	Endpoint    string
	UserAgent   string
	TokenSource TokenSource
}

func (c *RDTClient) CreateRestrictedDataToken(ctx context.Context, resources []RestrictedResource) (string, error) {
	if c.HTTP == nil {
		c.HTTP = http.DefaultClient
	}
	if c.Endpoint == "" {
		return "", fmt.Errorf("amazon sp-api: endpoint is required")
	}
	if c.TokenSource == nil {
		return "", fmt.Errorf("amazon sp-api: token source is required")
	}
	accessToken, err := c.TokenSource.AccessToken(ctx)
	if err != nil {
		return "", err
	}
	payload := struct {
		RestrictedResources []RestrictedResource `json:"restrictedResources"`
	}{RestrictedResources: resources}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(c.Endpoint, "/")+TokensPath, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("x-amz-access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode/100 != 2 {
		return "", fmt.Errorf("amazon sp-api: RDT request failed: status=%d request_id=%s body=%s", resp.StatusCode, resp.Header.Get("x-amzn-RequestId"), strings.TrimSpace(string(responseBody)))
	}
	var response struct {
		RestrictedDataToken string `json:"restrictedDataToken"`
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", err
	}
	if response.RestrictedDataToken == "" {
		return "", fmt.Errorf("amazon sp-api: RDT response has no restrictedDataToken")
	}
	return response.RestrictedDataToken, nil
}
