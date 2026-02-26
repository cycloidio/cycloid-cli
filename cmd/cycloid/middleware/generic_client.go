package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"reflect"
)

// GenericRequest add the default headers and api URL from cy context.
// You need to specify the org for authentication.
// The body and the response parameters must be pointers to struct.
func (m *middleware) GenericRequest(method string, org *string, params url.Values, headers map[string]string, body any, response any, route ...string) (*http.Response, error) {
	url, err := url.Parse(m.api.Config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url from %q, this means that CY_API_URL is probably invalid: %w", m.api.Config.URL, err)
	}

	url.Path = path.Join(route...)
	url.RawQuery = params.Encode()

	var bodyBytes []byte
	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize body to JSON for HTTP request: %w", err)
		}
	}

	req, err := http.NewRequest(method, url.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request with method %q and url %q: %w", method, url.String(), err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.Header.Add("Authorization", "Bearer "+m.api.GetToken(org))
	req.Header.Add("Content-Type", "application/json")

	resp, err := m.GenericClient.Do(req)
	if err != nil {
		return nil, NewAPIError(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body from HTTP request: %w", err)
	}

	responseReflectValue := reflect.ValueOf(response)
	if !responseReflectValue.IsNil() {

		if responseReflectValue.Kind() != reflect.Pointer {
			return nil, fmt.Errorf("GenericRequest response parameter requires a pointer to struct, this is an internal error, please report to the maintainer")
		}
		err = json.Unmarshal(respBody, response)
		if err != nil {
			return nil, fmt.Errorf("failed to decode JSON from body: %w", err)
		}
	}

	return resp, nil
}
