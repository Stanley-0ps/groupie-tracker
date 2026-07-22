package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const maxResponseSize = 10 << 20 // 10 MiB

var client = &http.Client{Timeout: 10 * time.Second}

// FetchJSON sends a GET request to url and decodes a single JSON value into
// target. The request is bound to ctx so it is cancelled when the caller goes
// away.
func FetchJSON(ctx context.Context, url string, target any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("GET %s: unexpected status %s", url, resp.Status)
	}

	decoder := json.NewDecoder(io.LimitReader(resp.Body, maxResponseSize))
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("decode response from %s: %w", url, err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return fmt.Errorf("decode response from %s: multiple JSON values", url)
		}
		return fmt.Errorf("decode response from %s: %w", url, err)
	}

	return nil
}
