package offline_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

// TestGetComponentConfig_ExplicitVersionID verifies that when an explicit versionID is
// provided, GetComponentConfig sends exactly ?service_catalog_source_version_id=<id>
// without any extra resolution calls.
func TestGetComponentConfig_ExplicitVersionID(t *testing.T) {
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

	_, _, err := m.GetComponentConfig("org", "proj", "env", "comp", 42)
	require.NoError(t, err)
	assert.Equal(t, "service_catalog_source_version_id=42", capturedQuery)
	assert.Equal(t, 1, callCount, "explicit version ID must not trigger extra API calls")
}

// TestGetComponentConfig_AutoResolvesLatestVersion verifies that when versionID == 0,
// GetComponentConfig resolves the component's latest stack version and passes the
// resolved service_catalog_source_version_id to the config endpoint.
func TestGetComponentConfig_AutoResolvesLatestVersion(t *testing.T) {
	const (
		org       = "myorg"
		project   = "myproj"
		env       = "prod"
		component = "mycomp"
		stackRef  = "myorg:my-stack"
		catRepo   = "my-catalog-repo"
		branch    = "main"
		wantID    = uint32(99)
	)

	var capturedConfigQuery string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path

		writeJSON := func(v any) {
			_ = json.NewEncoder(w).Encode(map[string]any{"data": v})
		}

		switch {
		case strings.HasSuffix(path, "/components/"+component+"/config"):
			capturedConfigQuery = r.URL.RawQuery
			writeJSON(map[string]any{})

		case strings.HasSuffix(path, "/components/"+component):
			writeJSON(map[string]any{
				"service_catalog": map[string]any{
					"ref":                              stackRef,
					"service_catalog_source_canonical": catRepo,
				},
			})

		case strings.Contains(path, "service_catalogs") && strings.HasSuffix(path, "/versions"):
			writeJSON([]map[string]any{
				{"id": wantID, "type": "branch", "name": branch, "commit_hash": "abc123def456"},
			})

		case strings.Contains(path, "service_catalogs"):
			writeJSON(map[string]any{
				"ref":                              stackRef,
				"service_catalog_source_canonical": catRepo,
			})

		case strings.Contains(path, "service_catalog_sources"):
			writeJSON(map[string]any{
				"canonical": catRepo,
				"branch":    branch,
			})

		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig(org, project, env, component, 0)
	require.NoError(t, err)
	assert.Equal(t, "service_catalog_source_version_id=99", capturedConfigQuery,
		"auto-resolve must pass the resolved version ID (99) to the config endpoint")
}
