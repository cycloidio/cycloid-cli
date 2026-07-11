package mockserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Version IDs and names served by ComponentConfigServer.
const (
	BranchVersionID = uint32(99)
	TagVersionID    = uint32(77)
	BranchName      = "main"
	TagName         = "v1.2.3"
	StackRef        = "myorg:my-stack"
	CatalogRepo     = "my-catalog-repo"
)

// ComponentConfigServer returns a test server that mocks the full resolution chain
// for GetComponentConfig: GetComponent → GetStack → GetCatalogRepository →
// ListStackVersions → config GET. The capturedQuery pointer receives the raw query
// string of the final config request.
func ComponentConfigServer(t *testing.T, capturedQuery *string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
					"ref":                              StackRef,
					"service_catalog_source_canonical": CatalogRepo,
				},
			})

		case strings.Contains(path, "service_catalogs") && strings.HasSuffix(path, "/versions"):
			writeJSON([]map[string]any{
				{"id": BranchVersionID, "type": "branch", "name": BranchName, "commit_hash": "abc123"},
				{"id": TagVersionID, "type": "tag", "name": TagName, "commit_hash": "def456"},
			})

		case strings.Contains(path, "service_catalogs"):
			writeJSON(map[string]any{
				"ref":                              StackRef,
				"service_catalog_source_canonical": CatalogRepo,
			})

		case strings.Contains(path, "service_catalog_sources"):
			writeJSON(map[string]any{
				"canonical": CatalogRepo,
				"branch":    BranchName,
			})

		default:
			http.NotFound(w, r)
		}
	}))
}
