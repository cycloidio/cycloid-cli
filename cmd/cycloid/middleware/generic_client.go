package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/spf13/viper"
)

// GenericRequest sends an HTTP request to the Cycloid API.
// It adds default authentication headers and handles JSON marshaling/unmarshaling.
// On non-2xx responses, it returns an *APIResponseError.
// On 2xx responses, it unwraps the {"data": <response>} envelope into the response pointer.
// Pass nil as response to discard the response body.
func (m *middleware) GenericRequest(req Request, response any) (*http.Response, error) {
	baseURL, err := url.Parse(m.api.Config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url from %q, this means that CY_API_URL is probably invalid: %w", m.api.Config.URL, err)
	}

	// Build URL from route segments
	routeParts := make([]string, 0, len(req.Route)+1)
	routeParts = append(routeParts, baseURL.Path)
	routeParts = append(routeParts, req.Route...)
	baseURL.Path = path.Join(routeParts...)

	// Encode query parameters
	if req.Query != nil {
		qv, err := encodeQuery(req.Query)
		if err != nil {
			return nil, fmt.Errorf("failed to encode query params: %w", err)
		}
		baseURL.RawQuery = qv.Encode()
	}

	// Marshal body
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, err = json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize body to JSON for HTTP request: %w", err)
		}
	}

	httpReq, err := http.NewRequest(req.Method, baseURL.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with method %q and url %q: %w", req.Method, baseURL.String(), err)
	}

	// Set headers
	for k, v := range req.Headers {
		httpReq.Header.Add(k, v)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Set Accept header
	if req.Accept != nil {
		httpReq.Header.Set("Accept", *req.Accept)
	}

	// Set auth header
	if !req.NoAuth {
		httpReq.Header.Set("Authorization", "Bearer "+m.api.GetToken(req.Organization))
	}

	debug := viper.GetString("verbosity") == "debug"
	if debug {
		httpDebugLogger.logRequest(httpReq, bodyBytes)
	}

	start := time.Now()
	resp, err := m.GenericClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	elapsed := time.Since(start)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body from HTTP response: %w", err)
	}

	if debug {
		httpDebugLogger.logResponse(resp, respBody, elapsed)
	}

	// Handle non-2xx responses
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, newAPIResponseError(resp, respBody)
	}

	// Unwrap {"data": ...} envelope
	if response != nil && len(respBody) > 0 {
		var envelope struct {
			Data json.RawMessage `json:"data"`
		}
		if err := json.Unmarshal(respBody, &envelope); err != nil {
			// Try direct unmarshal if envelope fails
			if err2 := json.Unmarshal(respBody, response); err2 != nil {
				return resp, fmt.Errorf("failed to decode JSON response: %w (envelope error: %v)", err2, err)
			}
			return resp, nil
		}

		if envelope.Data != nil {
			if err := json.Unmarshal(envelope.Data, response); err != nil {
				return resp, fmt.Errorf("failed to decode JSON from data envelope: %w", err)
			}
		}
	}

	return resp, nil
}
