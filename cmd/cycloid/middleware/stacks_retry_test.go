package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func withStackVersionResolveRetryDelays(t *testing.T, delays []time.Duration) {
	t.Helper()

	originalDelays := stackVersionResolveRetryDelays
	stackVersionResolveRetryDelays = delays
	t.Cleanup(func() {
		stackVersionResolveRetryDelays = originalDelays
	})
}

func newStackVersionTestMiddleware(t *testing.T, handler http.HandlerFunc) *middleware {
	t.Helper()

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	m, ok := NewMiddleware(common.NewAPI(common.WithURL(server.URL), common.WithToken("token"))).(*middleware)
	require.True(t, ok)
	return m
}

func TestResolveStackVersionRetriesBranchNotFound(t *testing.T) {
	withStackVersionResolveRetryDelays(t, []time.Duration{0, 0, 0})

	var calls int32
	m := newStackVersionTestMiddleware(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/organizations/org/service_catalogs/org:stack/versions", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		attempt := atomic.AddInt32(&calls, 1)
		if attempt < 3 {
			fmt.Fprint(w, `{"data":[{"id":1,"commit_hash":"1111111","name":"main","type":"branch"}]}`)
			return
		}

		fmt.Fprint(w, `{"data":[{"id":42,"commit_hash":"abc1234","name":"feature/my-feature-branch","type":"branch"}]}`)
	})

	versionID, commitHash, err := m.resolveStackVersion("org", "org:stack", "", "feature/my-feature-branch", "")

	require.NoError(t, err)
	assert.Equal(t, uint32(42), versionID)
	assert.Equal(t, "abc1234", commitHash)
	assert.Equal(t, int32(3), atomic.LoadInt32(&calls))
}

func TestResolveStackVersionReturnsNotFoundAfterRetries(t *testing.T) {
	withStackVersionResolveRetryDelays(t, []time.Duration{0, 0})

	var calls int32
	m := newStackVersionTestMiddleware(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		atomic.AddInt32(&calls, 1)
		fmt.Fprint(w, `{"data":[{"id":1,"commit_hash":"1111111","name":"main","type":"branch"}]}`)
	})

	_, _, err := m.resolveStackVersion("org", "org:stack", "", "feature/my-feature-branch", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), `stack version branch "feature/my-feature-branch" not found`)
	assert.Equal(t, int32(3), atomic.LoadInt32(&calls))
}

func TestResolveStackVersionDoesNotRetryListErrors(t *testing.T) {
	withStackVersionResolveRetryDelays(t, []time.Duration{0, 0})

	var calls int32
	m := newStackVersionTestMiddleware(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		http.Error(w, "backend unavailable", http.StatusInternalServerError)
	})

	_, _, err := m.resolveStackVersion("org", "org:stack", "", "feature/my-feature-branch", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list stack versions")
	assert.Equal(t, int32(1), atomic.LoadInt32(&calls))
}
