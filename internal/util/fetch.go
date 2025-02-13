package util

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const defaultTimeout = 2 * time.Second

type GetOpts struct {
	Timeout time.Duration
	URL     string
}

func GetWithTimeout(ctx context.Context, responseObj any, opts GetOpts) error {
	if IsNilPointer(responseObj) {
		return fmt.Errorf("responseObj must be a non-nil pointer")
	}

	if opts.URL == "" {
		return fmt.Errorf("URL is required")
	}

	timeout := opts.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, opts.URL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(responseObj); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}