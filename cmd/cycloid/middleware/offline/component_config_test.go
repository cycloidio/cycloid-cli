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

// componentConfigServer builds a test server that mocks the resolution chain for
// GetComponentConfig. It returns the stack version whose (type, name) matches
// wantType/wantName for explicit resolution, or returns the branch head when
// catalogRepoBranch matches a branch-type version (default/auto-resolve path).
func componentConfigServer(t *testing.T, capturedQuery *string) *httptest.Server {
	t.Helper()
	const (
		stackRef   = "myorg:my-stack"
		catRepo    = "my-catalog-repo"
		branchName = "main"
		branchID   = uint32(99)
		tagName    = "v1.2.3"
		tagID      = uint32(77)
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path

		writeJSON := func(v any) {
			_ = json.NewEncoder(w).Encode(map[string]any{"data": v})
		}

		switch {
		case strings.HasSuffix(path, "/components/mycomp/config"):
			*capturedQuery = r.URL.RawQuery
			writeJSON(map[string]any{})

		case strings.HasSuffix(path, "/components/mycomp"):
			writeJSON(map[string]any{
				"service_catalog": map[string]any{
					"ref":                              stackRef,
					"service_catalog_source_canonical": catRepo,
				},
			})

		case strings.Contains(path, "service_catalogs") && strings.HasSuffix(path, "/versions"):
			writeJSON([]map[string]any{
				{"id": branchID, "type": "branch", "name": branchName, "commit_hash": "abc123"},
				{"id": tagID, "type": "tag", "name": tagName, "commit_hash": "def456"},
			})

		case strings.Contains(path, "service_catalogs"):
			writeJSON(map[string]any{
				"ref":                              stackRef,
				"service_catalog_source_canonical": catRepo,
			})

		case strings.Contains(path, "service_catalog_sources"):
			writeJSON(map[string]any{
				"canonical": catRepo,
				"branch":    branchName,
			})

		default:
			http.NotFound(w, r)
		}
	}))
	return srv
}

// TestGetComponentConfig_BranchResolvesVersion verifies that passing an explicit
// branch name resolves to that branch's version ID in the config request.
func TestGetComponentConfig_BranchResolvesVersion(t *testing.T) {
	var capturedQuery string
	srv := componentConfigServer(t, &capturedQuery)
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig("myorg", "myproj", "myenv", "mycomp", "", "main", "")
	require.NoError(t, err)
	assert.Equal(t, "service_catalog_source_version_id=99", capturedQuery,
		"branch 'main' should resolve to ID 99")
}

// TestGetComponentConfig_TagResolvesVersion verifies that passing a tag name
// resolves to that tag's version ID in the config request.
func TestGetComponentConfig_TagResolvesVersion(t *testing.T) {
	var capturedQuery string
	srv := componentConfigServer(t, &capturedQuery)
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig("myorg", "myproj", "myenv", "mycomp", "v1.2.3", "", "")
	require.NoError(t, err)
	assert.Equal(t, "service_catalog_source_version_id=77", capturedQuery,
		"tag 'v1.2.3' should resolve to ID 77")
}

// TestGetComponentConfig_AutoResolvesLatestVersion verifies that when no version
// is specified, GetComponentConfig resolves the catalog-repo branch head and
// passes the corresponding service_catalog_source_version_id to the config endpoint.
func TestGetComponentConfig_AutoResolvesLatestVersion(t *testing.T) {
	var capturedQuery string
	srv := componentConfigServer(t, &capturedQuery)
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	_, _, err := m.GetComponentConfig("myorg", "myproj", "myenv", "mycomp", "", "", "")
	require.NoError(t, err)
	assert.Equal(t, "service_catalog_source_version_id=99", capturedQuery,
		"default (no version specified) should resolve to catalog branch head ID 99")
}
