package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd"
)

// Route defines a canned HTTP response for a given method+path pair.
type Route struct {
	Method     string
	Path       string
	Status     int
	Response   interface{}
	HandleFunc http.HandlerFunc
}

// CannedServer returns an httptest.Server that replies with the canned
// responses defined by routes. If a request matches no route the server
// returns 404. The server is automatically closed when the test finishes.
func CannedServer(t *testing.T, routes []Route) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()

	for _, r := range routes {
		route := r
		pattern := route.Method + " " + route.Path
		mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
			if route.HandleFunc != nil {
				route.HandleFunc(w, req)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			status := route.Status
			if status == 0 {
				status = http.StatusOK
			}
			w.WriteHeader(status)
			if route.Response != nil {
				if err := json.NewEncoder(w).Encode(route.Response); err != nil {
					t.Errorf("CannedServer: failed to encode response: %v", err)
				}
			}
		})
	}

	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

// CLIResult holds the captured output from a CLI command execution.
type CLIResult struct {
	Stdout string
	Stderr string
	Err    error
}

// RunCLI executes the CLI root command with the given arguments, pointing
// the API URL at serverURL and injecting a dummy API key + org.
// It captures stdout and returns the result.
func RunCLI(t *testing.T, serverURL string, args ...string) CLIResult {
	t.Helper()
	return RunCLIWithSetup(t, serverURL, nil, args...)
}

// RunCLIWithSetup is like RunCLI but accepts a setup function that can
// modify the root command before execution.
func RunCLIWithSetup(t *testing.T, serverURL string, setup func(root *cobra.Command), args ...string) CLIResult {
	t.Helper()

	// Configure viper and env vars to point at the test server.
	// Save and restore viper state so tests don't leak global config.
	prevURL := viper.GetString("api-url")
	prevOrg := viper.GetString("org")
	viper.Set("api-url", serverURL)
	if prevOrg == "" {
		viper.Set("org", "test-org")
	}
	t.Setenv("CY_API_KEY", "test-api-key")
	t.Cleanup(func() {
		viper.Set("api-url", prevURL)
		viper.Set("org", prevOrg)
	})

	root := cmd.NewRootCommand()
	if setup != nil {
		setup(root)
	}

	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	root.SetArgs(append([]string{"--output", "json"}, args...))

	err := root.Execute()

	return CLIResult{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
		Err:    err,
	}
}
