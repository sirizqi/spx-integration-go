package utils

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	},
}

// Logging
func Do(ctx context.Context, method, url string, body []byte, headers map[string]string) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return nil, 0, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Before Req
	LogOutgoingRequest(method, url, headers, body)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	// After Resp
	LogIncomingResponse(resp, respBody)

	return respBody, resp.StatusCode, nil
}
