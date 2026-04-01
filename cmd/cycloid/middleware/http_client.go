package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// StackVersion is the local representation of a stack catalog version.
// The model was removed from the generated swagger client.
type StackVersion struct {
	CommitHash *string `json:"commit_hash"`
	ID         *uint32 `json:"id"`
	IsLatest   *bool   `json:"is_latest"`
	Name       *string `json:"name"`
	Status     *string `json:"status"`
	Type       *string `json:"type"`
	Usage      *int64  `json:"usage"`
}

// StackUseCase is the local representation of a stack use case.
// The model was removed from the generated swagger client.
type StackUseCase struct {
	CloudProvider string  `json:"cloud_provider,omitempty"`
	Description   string  `json:"description,omitempty"`
	Name          *string `json:"name"`
	UseCase       *string `json:"use_case"`
}

// Request represents an HTTP request to the Cycloid API.
type Request struct {
	Method       string
	Organization *string  // used for auth token lookup
	NoAuth       bool     // disables auth header
	Route        []string // joined onto base URL: ["organizations", org, "projects"]
	Query        any      // url.Values or struct with `url` tags
	Headers      map[string]string
	Accept       *string // overrides default Accept header
	Body         any     // JSON-marshalled when non-nil
}

// APIResponseError is returned when the API returns a non-2xx response.
type APIResponseError struct {
	StatusCode int
	Status     string
	Body       []byte
	Payload    *models.ErrorPayload
	Path       string
}

func (e *APIResponseError) Error() string {
	if e.Payload != nil && len(e.Payload.Errors) > 0 && e.Payload.Errors[0].Message != nil {
		return fmt.Sprintf("API error %d: %s", e.StatusCode, *e.Payload.Errors[0].Message)
	}

	body := strings.TrimSpace(string(e.Body))
	if body != "" {
		if e.Path != "" {
			return fmt.Sprintf("API error %d on %q: %s", e.StatusCode, e.Path, body)
		}
		return fmt.Sprintf("API error %d: %s", e.StatusCode, body)
	}

	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Status)
}

// GetPayload implements ErrorPayloader for backwards compatibility.
func (e *APIResponseError) GetPayload() *models.ErrorPayload {
	return e.Payload
}

// HTTPStatusCode implements printer.ErrHTTPResponse.
func (e *APIResponseError) HTTPStatusCode() int {
	return e.StatusCode
}

// HTTPResponseBody implements printer.ErrHTTPResponse.
func (e *APIResponseError) HTTPResponseBody() []byte {
	return e.Body
}

// HTTPRequestPath implements printer.RequestPather.
func (e *APIResponseError) HTTPRequestPath() string {
	return e.Path
}

func newAPIResponseError(resp *http.Response, body []byte) *APIResponseError {
	apiErr := &APIResponseError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       body,
	}
	if resp.Request != nil && resp.Request.URL != nil {
		apiErr.Path = resp.Request.URL.RequestURI()
	}

	var payload models.ErrorPayload
	if err := json.Unmarshal(body, &payload); err == nil && len(payload.Errors) > 0 {
		apiErr.Payload = &payload
	}

	return apiErr
}
