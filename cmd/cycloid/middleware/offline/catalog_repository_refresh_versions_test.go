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

// catalogRepoResponse returns a minimal valid service_catalog_source JSON response.
func catalogRepoResponse(canonical string) string {
	return fmt.Sprintf(`{"data":{"canonical":%q,"name":%q,"url":"git@github.com:my/repo.git","branch":"main"}}`, canonical, canonical)
}

// versionsRefreshResponse returns a minimal versions/refresh JSON response.
func versionsRefreshResponse() string {
	return `{"data":[{"id":1,"commit_hash":"abc1234","name":"main","type":"branch"},{"id":2,"commit_hash":"def5678","name":"feat/my-branch","type":"branch"}]}`
}

// TestRefreshCatalogRepositoryVersions_CalledOnceWithFlag verifies that RefreshCatalogRepositoryVersions
// hits the versions/refresh endpoint exactly once when invoked.
func TestRefreshCatalogRepositoryVersions_CalledOnceWithFlag(t *testing.T) {
	var refreshCalls int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/organizations/org/service_catalog_sources/my-catalog/versions/refresh":
			atomic.AddInt32(&refreshCalls, 1)
			assert.Equal(t, "true", r.URL.Query().Get("sync_presence"), "sync_presence query param must be true")
			fmt.Fprint(w, versionsRefreshResponse())
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(srv.URL), common.WithToken("token")))

	versions, _, err := m.RefreshCatalogRepositoryVersions("org", "my-catalog")
	require.NoError(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&refreshCalls), "versions/refresh called exactly once")
	assert.Len(t, versions, 2)
}

// TestCreateCatalogRepository_RefreshCalledWhenFlagSet simulates the full create + refresh flow:
// POST creates the repo, then GET versions/refresh is called, confirming --refresh triggers exactly
// one refresh call and zero calls without the flag.
func TestCreateCatalogRepository_RefreshCalledWhenFlagSet(t *testing.T) {
	var createCalls int32
	var refreshCalls int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/organizations/org/service_catalog_sources":
			atomic.AddInt32(&createCalls, 1)
			// Verify body has required fields
			var body map[string]any
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			assert.NotEmpty(t, body["name"])
			fmt.Fprint(w, catalogRepoResponse("my-catalog"))
		case r.Method == http.MethodGet && r.URL.Path == "/organizations/org/service_catalog_sources/my-catalog/versions/refresh":
			atomic.AddInt32(&refreshCalls, 1)
			assert.Equal(t, "true", r.URL.Query().Get("sync_presence"), "sync_presence query param must be true")
			fmt.Fprint(w, versionsRefreshResponse())
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(srv.URL), common.WithToken("token")))

	// Simulate create (the command does this then calls RefreshCatalogRepositoryVersions when --refresh is set)
	cr, _, err := m.CreateCatalogRepository("org", "my-catalog", "git@github.com:my/repo.git", "main", "", "", "")
	require.NoError(t, err)
	require.NotNil(t, cr)
	assert.Equal(t, int32(1), atomic.LoadInt32(&createCalls))
	assert.Equal(t, int32(0), atomic.LoadInt32(&refreshCalls), "no refresh yet — flag not set")

	// Now simulate --refresh: the command calls RefreshCatalogRepositoryVersions after create
	_, _, err = m.RefreshCatalogRepositoryVersions("org", "my-catalog")
	require.NoError(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&refreshCalls), "refresh called exactly once with flag")
}

// TestUpdateCatalogRepository_RefreshCalledWhenFlagSet mirrors the create test for update.
func TestUpdateCatalogRepository_RefreshCalledWhenFlagSet(t *testing.T) {
	var updateCalls int32
	var refreshCalls int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPut && r.URL.Path == "/organizations/org/service_catalog_sources/my-catalog":
			atomic.AddInt32(&updateCalls, 1)
			fmt.Fprint(w, catalogRepoResponse("my-catalog"))
		case r.Method == http.MethodGet && r.URL.Path == "/organizations/org/service_catalog_sources/my-catalog/versions/refresh":
			atomic.AddInt32(&refreshCalls, 1)
			assert.Equal(t, "true", r.URL.Query().Get("sync_presence"), "sync_presence query param must be true")
			fmt.Fprint(w, versionsRefreshResponse())
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(srv.URL), common.WithToken("token")))

	// Simulate update without --refresh
	cr, _, err := m.UpdateCatalogRepository("org", "my-catalog", "my-catalog", "git@github.com:my/repo.git", "main", "", nil)
	require.NoError(t, err)
	require.NotNil(t, cr)
	assert.Equal(t, int32(1), atomic.LoadInt32(&updateCalls))
	assert.Equal(t, int32(0), atomic.LoadInt32(&refreshCalls), "no refresh — flag not set")

	// Now simulate --refresh: the command calls RefreshCatalogRepositoryVersions after update
	_, _, err = m.RefreshCatalogRepositoryVersions("org", "my-catalog")
	require.NoError(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&refreshCalls), "refresh called exactly once with flag")
}

// TestCreateCatalogRepository_NoRefreshWithoutFlag verifies that when --refresh is not set,
// the versions/refresh endpoint is NOT called (zero calls).
func TestCreateCatalogRepository_NoRefreshWithoutFlag(t *testing.T) {
	var refreshCalls int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/organizations/org/service_catalog_sources":
			fmt.Fprint(w, catalogRepoResponse("my-catalog"))
		case r.Method == http.MethodGet && r.URL.Path == "/organizations/org/service_catalog_sources/my-catalog/versions/refresh":
			atomic.AddInt32(&refreshCalls, 1)
			t.Error("RefreshCatalogRepositoryVersions must not be called when --refresh flag is absent")
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(srv.URL), common.WithToken("token")))

	// Create without --refresh: only the POST happens, no refresh
	cr, _, err := m.CreateCatalogRepository("org", "my-catalog", "git@github.com:my/repo.git", "main", "", "", "")
	require.NoError(t, err)
	require.NotNil(t, cr)
	assert.Equal(t, int32(0), atomic.LoadInt32(&refreshCalls), "zero refresh calls without --refresh flag")
}

// TestUpdateCatalogRepository_NoRefreshWithoutFlag mirrors the create test for update.
func TestUpdateCatalogRepository_NoRefreshWithoutFlag(t *testing.T) {
	var refreshCalls int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPut && r.URL.Path == "/organizations/org/service_catalog_sources/my-catalog":
			fmt.Fprint(w, catalogRepoResponse("my-catalog"))
		case r.Method == http.MethodGet && r.URL.Path == "/organizations/org/service_catalog_sources/my-catalog/versions/refresh":
			atomic.AddInt32(&refreshCalls, 1)
			t.Error("RefreshCatalogRepositoryVersions must not be called when --refresh flag is absent")
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(srv.URL), common.WithToken("token")))

	// Update without --refresh: only the PUT happens, no refresh
	cr, _, err := m.UpdateCatalogRepository("org", "my-catalog", "my-catalog", "git@github.com:my/repo.git", "main", "", nil)
	require.NoError(t, err)
	require.NotNil(t, cr)
	assert.Equal(t, int32(0), atomic.LoadInt32(&refreshCalls), "zero refresh calls without --refresh flag")
}

// TestRefreshCatalogRepositoryVersions_PropagatesError verifies that a non-2xx response
// from the versions/refresh endpoint is propagated as an error to the caller.
func TestRefreshCatalogRepositoryVersions_PropagatesError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"errors":[{"code":"ServerError","message":"backend error"}]}`, http.StatusInternalServerError)
	}))
	defer srv.Close()

	m := middleware.NewMiddleware(common.NewAPI(common.WithURL(srv.URL), common.WithToken("token")))

	versions, _, err := m.RefreshCatalogRepositoryVersions("org", "my-catalog")
	require.Error(t, err)
	assert.Nil(t, versions)
}
