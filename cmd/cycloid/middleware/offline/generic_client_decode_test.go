package offline_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
)

func TestGenericRequest_200InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`not-valid-json`))
	}))
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	var out map[string]any
	httpResp, err := m.GenericRequest(middleware.Request{
		Method: http.MethodGet,
		Route:  []string{"test"},
	}, &out)

	require.Error(t, err)
	require.NotNil(t, httpResp)
	assert.Equal(t, http.StatusOK, httpResp.StatusCode)

	var httpErr printer.ErrHTTPResponse
	require.ErrorAs(t, err, &httpErr)
	assert.Equal(t, http.StatusOK, httpErr.HTTPStatusCode())
	assert.Contains(t, string(httpErr.HTTPResponseBody()), "not-valid-json")
}

func TestGenericRequest_200EnvelopeDataTypeMismatch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":true}`))
	}))
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	var out map[string]any
	httpResp, err := m.GenericRequest(middleware.Request{
		Method: http.MethodGet,
		Route:  []string{"test"},
	}, &out)

	require.Error(t, err)
	require.NotNil(t, httpResp)
	assert.Equal(t, http.StatusOK, httpResp.StatusCode)

	var httpErr printer.ErrHTTPResponse
	require.ErrorAs(t, err, &httpErr)
	assert.Equal(t, `{"data":true}`, string(httpErr.HTTPResponseBody()))
}
