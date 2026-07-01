package credentials_test

// Regression test for CLI-123: `credential update` with --name only (no --canonical) was
// sending PUT to the collection endpoint /organizations/{org}/credentials → 405.
// After the fix, it resolves the canonical via ListCredentials and sends PUT to the item
// endpoint /organizations/{org}/credentials/{canonical}.

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/credentials"
)

func TestUpdateByNameResolvesCanonical_ItemRoute(t *testing.T) {
	const (
		org       = "test-org"
		credName  = "nm-testing-custom-secret"
		credCanon = "nm-testing-custom-secret"
	)

	type req struct{ method, path string }
	var captured []req

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = append(captured, req{r.Method, r.URL.Path})
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/credentials"):
			// ListCredentials — return one credential matching the test name
			_, _ = fmt.Fprintf(w, `{"data":[{"canonical":%q,"name":%q,"path":%q,"type":"custom"}]}`,
				credCanon, credName, credCanon)
		case r.Method == http.MethodPut:
			// UpdateCredential — minimal valid response
			_, _ = fmt.Fprintf(w, `{"data":{"canonical":%q,"name":%q,"path":%q,"type":"custom","description":""}}`,
				credCanon, credName, credCanon)
		default:
			http.Error(w, fmt.Sprintf("unexpected %s %s", r.Method, r.URL.Path), http.StatusInternalServerError)
		}
	}))
	defer srv.Close()

	// Inject test server URL, org, and token via global Viper / env (what NewAPI+GetOrg read).
	viper.Set("api-url", srv.URL)
	viper.Set("org", org)
	t.Setenv("CY_API_KEY", "test-token")
	t.Cleanup(func() {
		viper.Set("api-url", "")
		viper.Set("org", "")
	})

	updateCmd := credentials.NewUpdateCommand()
	updateCmd.SetArgs([]string{"custom", "--name", credName, "--field", "tesnmsecret2=toto"})
	err := updateCmd.Execute()
	require.NoError(t, err)

	// Assert ListCredentials was called for canonical resolution.
	listPath := fmt.Sprintf("/organizations/%s/credentials", org)
	var listSeen bool
	for _, r := range captured {
		if r.method == http.MethodGet && r.path == listPath {
			listSeen = true
		}
	}
	assert.True(t, listSeen, "expected GET %s (ListCredentials) to be called for name resolution", listPath)

	// Assert PUT targeted the ITEM route, not the collection — regression guard for CLI-123.
	itemPath := fmt.Sprintf("/organizations/%s/credentials/%s", org, credCanon)
	var putPath string
	for _, r := range captured {
		if r.method == http.MethodPut {
			putPath = r.path
			break
		}
	}
	require.NotEmpty(t, putPath, "expected a PUT request to be sent")
	assert.Equal(t, itemPath, putPath,
		"PUT must target item route (CLI-123 regression: was hitting collection %q)", listPath)
}
