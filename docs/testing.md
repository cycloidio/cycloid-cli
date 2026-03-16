# Testing

## Overview

There are two test layers:

| Layer | Location | Runs against |
|-------|----------|--------------|
| Middleware unit tests | `cmd/cycloid/middleware/*_test.go` | Real backend (via `testcfg`) |
| E2E tests | `e2e/*_test.go` | Real backend (via `testcfg`) |

Both layers require a running Cycloid backend. See [Running tests](#running-tests).

## Running tests

```bash
# Start the backend
make be-reset

# Run all tests
make test
# or
go test ./...

# Run a specific middleware test
go test ./cmd/cycloid/middleware/... -run TestGetProject

# Run a specific e2e test
go test ./e2e/... -run TestProjects

# Stop the backend when done
make be-stop
```

`make be-reset` is preferred for test runs because it starts from a clean backend state.

E2E tests are **not run in parallel** — the backend uses git under the hood and cannot handle concurrent writes.

## Environment variables

| Variable | Purpose |
|----------|---------|
| `CY_API_URL` | Backend URL (default: from `make be-reset` output) |
| `CY_API_KEY` | API key for authentication |
| `CY_TEST_ROOT_ORG` | Root org canonical for provisioning |
| `API_LICENCE_KEY` | Licence key (required by some backend features) |
| `CY_TEST_PROVISION_API` | Set to `true` to provision test fixtures via API |
| `CY_TEST_VERBOSITY` | Set to `debug` to enable HTTP debug logs |

`CY_TEST_VERBOSITY=debug` propagates to `CY_VERBOSITY=debug` inside `TestMain`, enabling full request/response logging. Logs include credentials (redacted to last 5 chars of the token).

## `testcfg` package

`internal/testcfg` manages test fixture provisioning. It creates a real org, project, environment, and component against the backend.

### `NewConfig(name string)`

```go
config, err := testcfg.NewConfig("middleware")
defer config.Cleanup()
```

After a successful call, `config` exposes:

| Field | Type | Description |
|-------|------|-------------|
| `config.Org` | `string` | Organization canonical |
| `config.APIUrl` | `string` | Backend URL |
| `config.APIKey` | `string` | API key scoped to this test run |
| `config.Project` | `*models.Project` | Pre-created project |
| `config.Environment` | `*models.Environment` | Pre-created environment |
| `config.Component` | `*models.Component` | Pre-created component |
| `config.ConfigRepo` | `*models.ExternalBackend` | Pre-created config repository |
| `config.CatalogRepo` | `*models.ServiceCatalogSource` | Pre-created catalog repository |

### Cleanup

`config.Cleanup()` deletes all provisioned resources. Always defer it immediately after `NewConfig`:

```go
config, err := testcfg.NewConfig("middleware")
defer config.Cleanup()
```

### `AppendCleanup`

Register additional teardown for resources created inside a test:

```go
config.AppendCleanup(func() {
    m.DeleteProject(config.Org, projectCanonical)
})
```

### `RandomCanonical`

```go
canonical := testcfg.RandomCanonical("prefix")
// returns something like "prefix-a1b2c3"
```

Use this for test resource names to avoid conflicts across parallel test runs (different packages).

## Middleware unit tests

### Setup

Each test package has a `TestMain` in `middleware_test.go`:

```go
var config *testcfg.Config

func runMain(ctx context.Context, main *testing.M) (int, error) {
    var err error
    config, err = testcfg.NewConfig("middleware")
    defer config.Cleanup()
    if err != nil {
        return 1, fmt.Errorf("config setup failed: %w", err)
    }

    os.Setenv("CY_API_URL", config.APIUrl)
    os.Setenv("CY_API_KEY", config.APIKey)
    os.Setenv("CY_ORG", config.Org)

    if v := os.Getenv("CY_TEST_VERBOSITY"); v != "" {
        viper.Set("verbosity", v)
    }
    return main.Run(), nil
}
```

### Writing a middleware test

```go
// cmd/cycloid/middleware/organization_projects_test.go
package middleware_test

import (
    "testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProject(t *testing.T) {
    m := NewTestMiddleware()   // helper from middleware_test package
    result, _, err := m.GetProject(config.Org, *config.Project.Canonical)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, *config.Project.Canonical, *result.Canonical)
}
```

Focus on the happy path. Error paths are implicitly covered by e2e tests.

## E2E tests

### Setup

`e2e/e2e_test.go` has its own `TestMain` with the same `testcfg.NewConfig` pattern. Fixtures from `testcfg` are available via the package-level `config` variable.

### Helpers

```go
// Execute a CLI command and return stdout, error
cmdOut, cmdErr := executeCommand([]string{"--org", config.Org, "project", "list"})

// Execute with stdin input
cmdOut, cmdErr := executeCommandStdin(stdin, []string{"--org", config.Org, "pipeline", "update"})

// Write a temporary file
WriteFile(path, content)

// Random canonical
canonical := randomCanonical("prefix")  // local alias for testcfg.RandomCanonical
```

### Pattern: create → test → deferred delete

```go
import (
    "encoding/json"
    "testing"

    "github.com/cycloidio/cycloid-cli/client/models"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestWidgets(t *testing.T) {
    canonical := testcfg.RandomCanonical("test-widget")

    t.Run("SuccessWidgetCreate", func(t *testing.T) {
        cmdOut, cmdErr := executeCommand([]string{
            "--output", "json", "--org", config.Org,
            "widgets", "create", "--name", canonical,
        })
        require.NoError(t, cmdErr)

        defer t.Run("SuccessWidgetDelete", func(t *testing.T) {
            _, err := executeCommand([]string{
                "--output", "json", "--org", config.Org,
                "widgets", "delete", "--widget", canonical,
            })
            require.NoError(t, err)
        })

        var widget models.Widget
        err := json.Unmarshal([]byte(cmdOut), &widget)
        require.NoError(t, err)
        require.NotNil(t, widget.Canonical)
        assert.Equal(t, canonical, *widget.Canonical)

        t.Run("SuccessWidgetGet", func(t *testing.T) { ... })
        t.Run("SuccessWidgetUpdate", func(t *testing.T) { ... })
    })
}
```

The `defer` on `SuccessWidgetDelete` ensures cleanup even if inner tests fail.

### Known skipped tests

| Test | File | Reason |
|------|------|--------|
| `SuccessOrganizationsList` | `organizations_test.go` | Requires user token; e2e uses API key (422) |
| `TestInfraPolicies` | `infra_policies_test.go` | Skipped pending backend availability |
| `PipelineClearTaskCacheOk` | `pipelines_test.go` | Pending backend support |
| `TestConfigRepositories` (some subtests) | `config_repositories_test.go` | See BE-981 |

### JSON helpers

```go
// Extract field values from a JSON array by matching another field
ids, err := JSONListExtractFields(listOut, "id", "email", "^admin@cycloid.io$")
// Returns []string of "id" values where "email" matches the regex
```

## Fixtures

Shared test fixtures are defined in `e2e/helpers_test.go`:

- `TestPipelineSample` — minimal Concourse pipeline YAML
- `TestPipelineVariables` — matching pipeline vars YAML
- `TestInfraPolicySample` — Rego policy
- `TestTerraformPlanSample` — minimal Terraform JSON plan (used by terracost + infra policies)
