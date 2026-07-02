package offline_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware/offline/mockserver"
)

func TestGetComponentConfig_RawVersionID(t *testing.T) {
	var capturedQuery string
	callCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/config") {
			capturedQuery = r.URL.RawQuery
		}
		_, _ = w.Write([]byte(`{"data":{}}`))
	}))
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig("org", "proj", "env", "comp", "", "", "", 14)
	require.NoError(t, err)
	assert.Equal(t, "service_catalog_source_version_id=14", capturedQuery)
	assert.Equal(t, 1, callCount, "raw version ID must not trigger extra API calls")
}

func TestGetComponentConfig_BranchResolvesVersion(t *testing.T) {
	var capturedQuery string
	srv := mockserver.ComponentConfigServer(t, &capturedQuery)
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig("myorg", "myproj", "myenv", "mycomp", "", mockserver.BranchName, "", 0)
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("service_catalog_source_version_id=%d", mockserver.BranchVersionID), capturedQuery)
}

func TestGetComponentConfig_TagResolvesVersion(t *testing.T) {
	var capturedQuery string
	srv := mockserver.ComponentConfigServer(t, &capturedQuery)
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig("myorg", "myproj", "myenv", "mycomp", mockserver.TagName, "", "", 0)
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("service_catalog_source_version_id=%d", mockserver.TagVersionID), capturedQuery)
}

func TestGetComponentConfig_AutoResolvesLatestVersion(t *testing.T) {
	var capturedQuery string
	srv := mockserver.ComponentConfigServer(t, &capturedQuery)
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig("myorg", "myproj", "myenv", "mycomp", "", "", "", 0)
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("service_catalog_source_version_id=%d", mockserver.BranchVersionID), capturedQuery)
}
