# Adding a Command

This walkthrough adds a hypothetical `GET /organizations/{org}/widgets/{widget}` endpoint as a concrete example.

## 1. Find the endpoint

Check the Cycloid API reference at https://docs.cycloid.io/api/index.html. Note:

- The HTTP method and path
- Required vs optional parameters
- The response model name (e.g., `Widget`)

## 2. Ensure the model exists

Models live in `client/models/` — auto-generated from `swagger.yaml`. If the model already exists, skip this step.

If `swagger.yaml` was updated:

```bash
make client-generate   # regenerates client/models/ and client/ from swagger.yaml
```

Never edit files in `client/models/` by hand.

## 3. Add middleware interface methods (prefer idempotent flows)

Open `cmd/cycloid/middleware/middleware.go` and add the required signatures to the `Middleware` interface.
When the resource can be managed idempotently, also expose a `CreateOrUpdateX` helper:

```go
// cmd/cycloid/middleware/middleware.go
GetWidget(org, widget string) (*models.Widget, *http.Response, error)
CreateWidget(org string, newWidget *models.NewWidget) (*models.Widget, *http.Response, error)
UpdateWidget(org, widget string, update *models.UpdateWidget) (*models.Widget, *http.Response, error)
CreateOrUpdateWidget(org string, widget string, payload *models.NewWidget) (*models.Widget, *http.Response, error)
```

Return type conventions:
- Get/Create → `(*models.X, *http.Response, error)`
- List → `([]*models.X, *http.Response, error)`
- Delete/void → `(*http.Response, error)`

## 4. Implement middleware methods

Create or extend the feature file. Keep the category name aligned with the API docs tag/category.
For an org-scoped resource in the `Organization` category, use `organization_widgets.go`:

```go
// cmd/cycloid/middleware/organization_widgets.go
package middleware

import (
    "net/http"
    "github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetWidget(org, widget string) (*models.Widget, *http.Response, error) {
    var result *models.Widget
    resp, err := m.GenericRequest(Request{
        Method:       "GET",
        Organization: &org,
        Route:        []string{"organizations", org, "widgets", widget},
    }, &result)
    if err != nil {
        return nil, resp, err
    }
    return result, resp, nil
}

// Prefer idempotent flows when possible: detect state, then create or update.
func (m *middleware) CreateOrUpdateWidget(org, widget string, payload *models.NewWidget) (*models.Widget, *http.Response, error) {
    current, resp, err := m.GetWidget(org, widget)
    if err == nil && current != nil {
        update := &models.UpdateWidget{
            Name: payload.Name,
        }
        return m.UpdateWidget(org, widget, update)
    }
    if resp != nil && resp.StatusCode == http.StatusNotFound {
        return m.CreateWidget(org, payload)
    }
    return nil, resp, err
}
```

Key points:
- `Organization: &org` enables auth token lookup
- `Route` segments are path-joined onto `CY_API_URL`
- `GenericRequest` unwraps `{"data": ...}` automatically; pass `&result` directly
- Return `nil, resp, err` on error so callers can inspect the HTTP response
- If no dedicated idempotent API route exists, implement idempotency in middleware (`Get`/`List` then `Create` or `Update`)

## 5. Follow the `cmd/cycloid/teams` command pattern

`cmd/cycloid/teams` is the reference pattern for subcommands. The important logic:

- `cmd.go` defines the category command and wires all subcommands in one place (`AddCommand(...)`).
- Subcommands explicitly define `Args` (`cobra.NoArgs`, `cobra.MinimumNArgs(1)`, etc.) for consistent completion and UX.
- `get` and `delete` support multiple identifiers via positional args and return either one resource or a list.
- `create` supports idempotent behavior with `--update` and can update an existing resource when requested.
- Shared flag/completion logic lives in `internal/cyargs`; command files consume `cyargs.Add*Flag` / `cyargs.Get*`.
- Nested domains are modeled with nested commands (e.g. `teams members ...`).

## 6. Add cobra command files

Use a command directory name that matches the API docs category/tag naming.

```
cmd/cycloid/widgets/
  cmd.go     ← registers subcommands, called from cmd/root.go
  get.go     ← accepts identifiers as args and/or --widget
  create.go  ← supports --update idempotent behavior
  update.go  ← can delegate to create --update flow
  delete.go  ← accepts one or many identifiers
```

```go
// cmd/cycloid/widgets/get.go
package widgets

import (
    "fmt"

    "github.com/cycloidio/cycloid-cli/client/models"
    "github.com/pkg/errors"
    "github.com/spf13/cobra"

    "github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
    "github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
    "github.com/cycloidio/cycloid-cli/internal/cyargs"
    "github.com/cycloidio/cycloid-cli/printer"
    "github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetWidget() *cobra.Command {
    cmd := &cobra.Command{
        Use:               "get [widget_identifiers...]",
        Short:             "Get one or more widgets",
        Args:              cobra.ArbitraryArgs,
        ValidArgsFunction: cyargs.CompleteWidget,
        Example: `
cy --org my-org widgets get my-widget
cy --org my-org widgets get my-widget another-widget
cy --org my-org widgets get --widget my-widget
`,
        RunE: getWidget,
    }
    cyargs.AddOrgFlag(cmd)
    cyargs.AddWidgetFlag(cmd)
    return cmd
}

func getWidget(cmd *cobra.Command, args []string) error {
    // Step 1: retrieve ALL flags before NewAPI/NewMiddleware
    org, err := cyargs.GetOrg(cmd)
    if err != nil {
        return err
    }

    widget, err := cyargs.GetWidget(cmd)
    if err != nil {
        return err
    }

    output, err := cyargs.GetOutput(cmd)
    if err != nil {
        return err
    }

    identifiers := append([]string{}, args...)
    if widget != "" {
        identifiers = append(identifiers, widget)
    }
    if len(identifiers) == 0 {
        return errors.New("provide at least one widget identifier (arg or --widget)")
    }

    // Step 2: build printer
    p, err := factory.GetPrinter(output)
    if err != nil {
        return err
    }

    // Step 3: build API + middleware
    api := common.NewAPI()
    m := middleware.NewMiddleware(api)

    // Step 4: call middleware
    if len(identifiers) == 1 {
        result, _, err := m.GetWidget(org, identifiers[0])
        if err != nil {
            return printer.SmartPrint(p, nil, err, "unable to get widget", printer.Options{}, cmd.OutOrStderr())
        }
        return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
    }

    result := make([]*models.Widget, 0, len(identifiers))
    for _, identifier := range identifiers {
        widget, _, err := m.GetWidget(org, identifier)
        if err != nil {
            return printer.SmartPrint(p, nil, err, fmt.Sprintf("unable to get widget %q", identifier), printer.Options{}, cmd.OutOrStderr())
        }
        result = append(result, widget)
    }
    return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
```

Global command directives to follow:

- `get`: accept canonical/identifier from args and/or flag; show both styles in `Example`; multiple args return `[]*models.Widget`.
- `create`/`update`: share logic for idempotency (`create --update`); `update` should also allow creation.
- `delete`: accept canonical/identifier from args and/or flag; support multiple values; print deleted identifier(s) only.
- Shared flags/completion must live in `internal/cyargs` (except simple bool flags like `--update`, or truly command-specific flags).

## 7. Register in `cmd.go`

```go
// cmd/cycloid/widgets/cmd.go
package widgets

import "github.com/spf13/cobra"

func NewWidgetsCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "widgets",
        Short: "Manage widgets",
        Args:  cobra.NoArgs,
    }
    cmd.AddCommand(NewGetWidget())
    // cmd.AddCommand(NewListWidgets())
    // cmd.AddCommand(NewCreateWidget())
    return cmd
}
```

Register in `cmd/root.go`:

```go
rootCmd.AddCommand(widgets.NewWidgetsCmd())
```

## 8. Add flags

**Shared flags** (used by multiple commands) belong in `internal/cyargs/`:

```go
// internal/cyargs/widgets.go
const widgetFlagName = "widget"

func AddWidgetFlag(cmd *cobra.Command) string {
    cmd.Flags().String(widgetFlagName, "", "Widget canonical or identifier")
    _ = cmd.RegisterFlagCompletionFunc(widgetFlagName, CompleteWidget)
    return widgetFlagName
}

func GetWidget(cmd *cobra.Command) (string, error) {
    return cmd.Flags().GetString(widgetFlagName)
}
```

**Feature-specific flags** (one command only) can be inline in the command file:

```go
cmd.Flags().String("format", "json", "Output format for widget data")
```

Register shared flags in the constructor:

```go
func NewGetWidget() *cobra.Command {
    cmd := &cobra.Command{...}
    cyargs.AddOrgFlag(cmd)
    widgetFlag := cyargs.AddWidgetFlag(cmd)
    _ = widgetFlag // Example: cmd.MarkFlagRequired(widgetFlag) when needed
    return cmd
}
```

## 9. Write a middleware unit test

```go
// cmd/cycloid/middleware/organization_widgets_test.go
package middleware_test

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestGetWidget(t *testing.T) {
    m := NewTestMiddleware() // helper from middleware_test package
    result, _, err := m.GetWidget(config.Org, "my-widget")
    require.NoError(t, err)
    assert.NotNil(t, result)
}
```

The `config` global is set up by `TestMain` in `middleware_test.go`. See `docs/testing.md`.

## 10. Write an e2e test

```go
// e2e/widgets_test.go
package e2e_test

import (
    "encoding/json"
    "testing"

    "github.com/cycloidio/cycloid-cli/client/models"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestWidgets(t *testing.T) {
    widgetCanonical := testcfg.RandomCanonical("test-widget")

    t.Run("SuccessWidgetCreate", func(t *testing.T) {
        cmdOut, cmdErr := executeCommand([]string{
            "--output", "json",
            "--org", config.Org,
            "widgets", "create",
            "--name", widgetCanonical,
        })
        require.NoError(t, cmdErr)
        defer t.Run("SuccessWidgetDelete", func(t *testing.T) {
            _, deleteErr := executeCommand([]string{
                "--output", "json", "--org", config.Org,
                "widgets", "delete", "--widget", widgetCanonical,
            })
            require.NoError(t, deleteErr)
        })

        var widget models.Widget
        err := json.Unmarshal([]byte(cmdOut), &widget)
        require.NoError(t, err)
        require.NotNil(t, widget.Canonical)
        assert.Equal(t, widgetCanonical, *widget.Canonical)

        t.Run("SuccessWidgetGet", func(t *testing.T) {
            cmdOut, cmdErr := executeCommand([]string{
                "--output", "json", "--org", config.Org,
                "widgets", "get", "--widget", widgetCanonical,
            })
            require.NoError(t, cmdErr)
            var got models.Widget
            err := json.Unmarshal([]byte(cmdOut), &got)
            require.NoError(t, err)
            require.NotNil(t, got.Canonical)
            assert.Equal(t, widgetCanonical, *got.Canonical)
        })
    })
}
```

Pattern: create → nested subtests → deferred delete.

## 11. Verify

```bash
make format           # fix formatting
make lint             # catch issues
go test ./cmd/cycloid/middleware/... -run TestGetWidget   # unit test
go test ./e2e/... -run TestWidgets                        # e2e test (requires make be-reset)
go build -o cy .      # confirm it builds
```

## File naming checklist

```
cmd/cycloid/<feature>/
  cmd.go          ← registers subcommands with cobra
  list.go         ← ListX command
  get.go          ← GetX command
  create.go       ← CreateX command
  update.go       ← UpdateX command
  delete.go       ← DeleteX command
  common.go       ← shared helpers (optional)
```
