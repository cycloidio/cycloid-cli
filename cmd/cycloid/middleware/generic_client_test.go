package middleware_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func TestGenericRequest_ListProjects(t *testing.T) {
	m := config.Middleware

	var projects []*models.Project
	httpResp, err := m.GenericRequest(middleware.Request{
		Method:       http.MethodGet,
		Organization: &config.Org,
		Route:        []string{"organizations", config.Org, "projects"},
	}, &projects)

	require.NoError(t, err)
	require.NotNil(t, httpResp)
	assert.Equal(t, http.StatusOK, httpResp.StatusCode)
	assert.NotNil(t, projects)
}

func TestGenericRequest_NotFound(t *testing.T) {
	m := config.Middleware

	var out any
	httpResp, err := m.GenericRequest(middleware.Request{
		Method:       http.MethodGet,
		Organization: &config.Org,
		Route:        []string{"organizations", config.Org, "projects", "this-project-does-not-exist-xyz"},
	}, &out)

	require.Error(t, err)
	require.NotNil(t, httpResp)
	assert.Equal(t, http.StatusNotFound, httpResp.StatusCode)

	var apiErr *middleware.APIResponseError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, http.StatusNotFound, apiErr.StatusCode)
	assert.NotEmpty(t, apiErr.Error())
}

func TestGenericRequest_NoAuth(t *testing.T) {
	m := config.Middleware

	var out any
	httpResp, err := m.GenericRequest(middleware.Request{
		Method: http.MethodGet,
		Route:  []string{"organizations", config.Org, "projects"},
		NoAuth: true,
	}, &out)

	require.Error(t, err)
	require.NotNil(t, httpResp)
	assert.Equal(t, http.StatusForbidden, httpResp.StatusCode)
}
