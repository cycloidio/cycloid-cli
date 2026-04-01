package middleware

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIResponseError_PathFallback(t *testing.T) {
	u, err := url.Parse("https://http-api.cycloid.io/organizations/org/projects/project?foo=bar")
	assert.NoError(t, err)

	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Status:     "400 Bad Request",
		Request: &http.Request{
			URL: u,
		},
	}

	apiErr := newAPIResponseError(resp, []byte("raw backend error"))
	assert.Equal(t, http.StatusBadRequest, apiErr.StatusCode)
	assert.Equal(t, "/organizations/org/projects/project?foo=bar", apiErr.Path)
	assert.Equal(t, "API error 400 on \"/organizations/org/projects/project?foo=bar\": raw backend error", apiErr.Error())
}
