package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	clicmd "github.com/cycloidio/cycloid-cli/cmd"
)

// Route defines a canned HTTP response for a given method+path.
type Route struct {
	Method     string
	Path       string
	Status     int
	Response   any
	HandleFunc http.HandlerFunc
}

// CannedServer creates an httptest.Server that serves canned JSON responses
// for the given routes. Unmatched routes return 404.
func CannedServer(t *testing.T, routes []Route) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	for _, r := range routes {
		r := r
		pattern := r.Method + " " + r.Path
		mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
			if r.HandleFunc != nil {
				r.HandleFunc(w, req)
				return
			}
			w.Header().Set("Content-Type", "application/vnd.cycloid.io.v1+json")
			w.WriteHeader(r.Status)
			if r.Response != nil {
				json.NewEncoder(w).Encode(r.Response)
			}
		})
	}

	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

// initViper re-applies the global viper configuration that root.go's init()
// normally sets up. This must be called after viper.Reset() to restore env-var
// mapping and defaults so that CY_* environment variables and flag bindings
// work correctly in tests.
func initViper() {
	viper.SetEnvPrefix("CY")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("console_url", "https://console.cycloid.io")
}

// RunCLI executes a CLI command against the given test server and returns
// its stdout output and error. It sets up a fresh root command with
// --api-url pointing at the test server and the provided args.
func RunCLI(t *testing.T, serverURL string, args ...string) (string, error) {
	t.Helper()

	// Reset viper to avoid state leaking between tests, then re-apply the
	// global configuration that root.go's init() normally provides.
	viper.Reset()
	initViper()

	rootCmd := clicmd.NewRootCommand()

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs(append([]string{"--api-url", serverURL}, args...))

	err := rootCmd.Execute()
	return stdout.String(), err
}

// RunCLIWithSetup is like RunCLI but allows customizing the root command
// before execution (e.g., to set flags or inject environment).
func RunCLIWithSetup(t *testing.T, serverURL string, setup func(cmd *cobra.Command), args ...string) (string, error) {
	t.Helper()

	viper.Reset()
	initViper()

	rootCmd := clicmd.NewRootCommand()

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs(append([]string{"--api-url", serverURL}, args...))

	if setup != nil {
		setup(rootCmd)
	}

	err := rootCmd.Execute()
	return stdout.String(), err
}
