package offline_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func TestCreateOrUpdateComponentRetriesTransientBranchVersion(t *testing.T) {
	var versionCalls int32
	var componentUpserts int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/organizations/org/service_catalogs/org:stack/versions":
			attempt := atomic.AddInt32(&versionCalls, 1)
			if attempt == 1 {
				fmt.Fprint(w, `{"data":[{"id":1,"commit_hash":"1111111","name":"main","type":"branch"}]}`)
				return
			}

			fmt.Fprint(w, `{"data":[{"id":42,"commit_hash":"abc1234","name":"feature/my-feature-branch","type":"branch"}]}`)

		case r.Method == http.MethodPut && r.URL.Path == "/organizations/org/projects/project/environments/prod/components":
			atomic.AddInt32(&componentUpserts, 1)

			var body map[string]any
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			assert.Equal(t, float64(42), body["service_catalog_source_version_id"])
			assert.Equal(t, "abc1234", body["service_catalog_source_version_commit_hash"])

			fmt.Fprint(w, `{"data":{"canonical":"mycompo"}}`)

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(server.URL), common.WithToken("token")))
	component, _, err := m.CreateOrUpdateComponent(
		"org",
		"project",
		"prod",
		"mycompo",
		"Component Description here.",
		"Name Here",
		"org:stack",
		"",
		"feature/my-feature-branch",
		"",
		"default",
		"",
		nil,
	)

	require.NoError(t, err)
	require.NotNil(t, component)
	assert.Equal(t, int32(2), atomic.LoadInt32(&versionCalls))
	assert.Equal(t, int32(1), atomic.LoadInt32(&componentUpserts))
}

func TestCreateOrUpdateComponentDoesNotRetryVersionListErrors(t *testing.T) {
	var versionCalls int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/organizations/org/service_catalogs/org:stack/versions":
			atomic.AddInt32(&versionCalls, 1)
			http.Error(w, "backend unavailable", http.StatusInternalServerError)
		case r.Method == http.MethodPut && r.URL.Path == "/organizations/org/projects/project/environments/prod/components":
			t.Fatal("component upsert should not be called when version listing fails")
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(server.URL), common.WithToken("token")))
	_, _, err := m.CreateOrUpdateComponent(
		"org",
		"project",
		"prod",
		"mycompo",
		"Component Description here.",
		"Name Here",
		"org:stack",
		"",
		"feature/my-feature-branch",
		"",
		"default",
		"",
		nil,
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list stack versions")
	assert.Equal(t, int32(1), atomic.LoadInt32(&versionCalls))
}
