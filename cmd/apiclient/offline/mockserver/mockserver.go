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
)

// ComponentConfigServer mocks the selector resolution chain of GetComponentConfig
// and GetComponentStackConfig (GetComponent → ListStackVersions → config GET) and
// stores the raw query string of the final config request in capturedQuery
func ComponentConfigServer(t *testing.T, capturedQuery *string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path

		writeJSON := func(v any) {
			_ = json.NewEncoder(w).Encode(map[string]any{"data": v})
		}

		switch {
		case strings.HasSuffix(path, "/components/mycomp/config"),
			strings.HasSuffix(path, "/components/mycomp/stack_config"):
			*capturedQuery = r.URL.RawQuery
			writeJSON(map[string]any{})

		case strings.HasSuffix(path, "/components/mycomp"):
			writeJSON(map[string]any{
				"service_catalog": map[string]any{
					"ref": StackRef,
				},
			})

		case strings.Contains(path, "service_catalogs") && strings.HasSuffix(path, "/versions"):
			writeJSON([]map[string]any{
				{"id": BranchVersionID, "type": "branch", "name": BranchName, "commit_hash": "abc123"},
				{"id": TagVersionID, "type": "tag", "name": TagName, "commit_hash": "def456"},
			})

		default:
			http.NotFound(w, r)
		}
	}))
}
