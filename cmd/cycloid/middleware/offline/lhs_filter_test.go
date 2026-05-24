package offline_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

// captureRawQuery runs a GenericRequest with the given LHSFilters against a local
// test server and returns the raw query string the server received.
func captureRawQuery(t *testing.T, filters []middleware.LHSFilter) string {
	t.Helper()
	var captured string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":[]}`))
	}))
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	var out []any
	_, err := m.GenericRequest(middleware.Request{
		Method:     http.MethodGet,
		Route:      []string{"test"},
		LHSFilters: filters,
	}, &out)
	require.NoError(t, err)
	return captured
}

func TestLHSFilterBracketsNotEncoded(t *testing.T) {
	raw := captureRawQuery(t, []middleware.LHSFilter{
		{Attribute: "name", Condition: "eq", Value: "my-project"},
	})
	assert.Equal(t, "name[eq]=my-project", raw, "brackets must be literal, not percent-encoded")
}

func TestLHSFilterMultipleFilters(t *testing.T) {
	raw := captureRawQuery(t, []middleware.LHSFilter{
		{Attribute: "name", Condition: "eq", Value: "my-project"},
		{Attribute: "canonical", Condition: "rlike", Value: "proj.*"},
	})
	assert.Equal(t, "name[eq]=my-project&canonical[rlike]=proj.*", raw)
}

func TestLHSFilterRegexMetacharsPreserved(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantValue string
	}{
		{"dot wildcard", "dg.growy", "dg.growy"},
		{"question mark", "dg.?growy", "dg.?growy"},
		{"star wildcard", "dg.*growy", "dg.*growy"},
		{"plus quantifier", "dg.+growy", "dg.+growy"},
		{"brackets", "[A-Z]+", "[A-Z]+"},
		{"pipe alternation", "foo|bar", "foo|bar"},
		{"caret anchor", "^foo", "^foo"},
		{"dollar anchor", "foo$", "foo$"},
		{"backslash escape", `foo\.bar`, `foo\.bar`},
		{"parens groups", "(foo|bar)", "(foo|bar)"},
		{"curly braces", "a{2,4}", "a{2,4}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := captureRawQuery(t, []middleware.LHSFilter{
				{Attribute: "name", Condition: "rlike", Value: tt.value},
			})
			assert.Equal(t, "name[rlike]="+tt.wantValue, raw)
		})
	}
}

func TestLHSFilterStructuralCharsEncoded(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantValue string
	}{
		{"ampersand", "foo&bar", "foo%26bar"},
		{"equals", "foo=bar", "foo%3Dbar"},
		{"hash", "foo#bar", "foo%23bar"},
		{"space", "foo bar", "foo%20bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := captureRawQuery(t, []middleware.LHSFilter{
				{Attribute: "name", Condition: "eq", Value: tt.value},
			})
			assert.Equal(t, "name[eq]="+tt.wantValue, raw)
		})
	}
}

func TestLHSFilterCombinedWithRegularQuery(t *testing.T) {
	var captured string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":[]}`))
	}))
	defer srv.Close()

	api := common.NewAPI(common.WithURL(srv.URL), common.WithToken("test-token"))
	m := middleware.NewMiddleware(api)

	type queryParams struct {
		PageSize int `url:"page_size"`
	}

	var out []any
	_, err := m.GenericRequest(middleware.Request{
		Method: http.MethodGet,
		Route:  []string{"test"},
		Query:  queryParams{PageSize: 100},
		LHSFilters: []middleware.LHSFilter{
			{Attribute: "name", Condition: "eq", Value: "my-project"},
		},
	}, &out)
	require.NoError(t, err)

	assert.Contains(t, captured, "page_size=100")
	assert.Contains(t, captured, "name[eq]=my-project")
}

func TestLHSFilterEmpty(t *testing.T) {
	raw := captureRawQuery(t, nil)
	assert.Equal(t, "", raw, "no filters should produce no query string")
}
