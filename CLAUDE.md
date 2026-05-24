# CLAUDE.md

This file is auto-loaded by Claude Code on every invocation. See also `AGENTS.md` (kept identical, for non-Claude LLM agents).

## Commands

```bash
# Build
make build                  # build all platform binaries
go build -o cy .            # quick local build

# Test (requires local backend)
make be-reset               # start backend via docker compose
go test ./...               # run all tests
make test                   # same
make be-stop                # stop backend

# Run a specific test
go test ./e2e/... -run TestProjects
go test ./cmd/cycloid/middleware/... -run TestGetProject

# Lint & format
make lint                   # golangci-lint + shellcheck
make format                 # gci + goimports + shfmt

# Client regeneration (when swagger.yaml changes)
make client-generate        # regenerates ./client/ from swagger.yaml

# Changelog
make new-changelog-entry    # add unreleased changelog entry (uses changie via docker)
```

## Architecture

```
cmd/cycloid/<feature>/*.go  →  cmd/cycloid/middleware/  →  GenericRequest()
     (cobra commands)              (Middleware interface)      (generic_client.go)
```

### Key packages

- **`client/models/`** — Auto-generated go-swagger models from `swagger.yaml`. Never edit manually. Regenerate with `make client-generate`.
- **`client/client/`** — Intentionally NOT used. The generated operations package was removed in the middleware refactor. Use `GenericRequest` instead.
- **`cmd/cycloid/middleware/`** — The `Middleware` interface (`middleware.go`) makes HTTP calls via `GenericRequest` (`generic_client.go`). Each feature gets its own file (e.g., `organization_projects.go`).
- **`cmd/cycloid/<feature>/`** — Cobra command definitions. Each feature directory has: `cmd.go` (registers subcommands), plus one file per verb (`list.go`, `get.go`, `create.go`, `update.go`, `delete.go`), and `common.go` for shared logic.
- **`internal/cyargs/`** — All shared flag definitions and completion functions. Every flag used by multiple commands must be declared here.
- **`printer/`** — Output formatting. Use `cyout.PrintWithOptions(cmd, obj, err, errMsg, opts)` — routes errors to `cmd.ErrOrStderr()`, results to `cmd.OutOrStdout()`. Direct `printer.SmartPrint` calls are legacy; migrate on touch.
- **`internal/cyout/`** — One-liner print helpers (`cyout.Print`, `cyout.PrintWithOptions`) and `cyout.RegisterModel` for `--output` shell completion.
- **`e2e/`** — End-to-end tests that run real CLI commands against a live backend.

## Hard Rules

These are invariants that LLM agents and new contributors must not violate:

1. **NEVER edit `client/models/`** — auto-generated. Run `make client-generate` to update.
2. **NEVER import or use `client/client/`** — deprecated. Use `GenericRequest` in middleware methods.
3. **NEVER call `NewMiddleware` outside a cobra `RunE` function** — not at package init, not in tests directly.
4. **ALWAYS parse ALL flags via `cyargs.Get*` BEFORE calling `common.NewAPI()` and `NewMiddleware()`** — `GenericRequest` reads `verbosity` from Viper at call time; unparsed flags produce stale values.
5. **ALWAYS add shared flags to `internal/cyargs/`** — never inline a flag used by more than one command.
6. **ALWAYS use `cyout.PrintWithOptions` (or `cyout.Print`)** — errors go to `ErrOrStderr()`, results go to `OutOrStdout()`. Do not call `printer.SmartPrint` directly in new code.
7. **E2E tests require a running backend** (`make be-start`). Never run e2e in parallel.
8. **Run `make format && make lint` after every code change.**
9. **Ship tests with every feature** — in the same change, add or extend coverage for what you introduce: new or changed middleware in `cmd/cycloid/middleware/*_test.go` (or focused unit tests where appropriate), and user-facing CLI behavior in `e2e/*_test.go` when that resource already has e2e tests. Do not land behavior-only changes without tests.

## `GenericRequest` pattern

`GenericRequest` handles auth, JSON marshaling, and `{"data":...}` envelope unwrapping. Pass `&result` directly — do not wrap in a `struct{ Data *X }`.

```go
// Standard middleware method pattern:
func (m *middleware) GetProject(org, project string) (*models.Project, *http.Response, error) {
    var result *models.Project
    resp, err := m.GenericRequest(Request{
        Method:       "GET",
        Organization: &org,
        Route:        []string{"organizations", org, "projects", project},
    }, &result)
    if err != nil {
        return nil, resp, err
    }
    return result, resp, nil
}
```

`Request` fields: `Method`, `Organization` (*string, for auth), `NoAuth` (bool), `Route` ([]string), `Query` (struct with `url` tags), `LHSFilters` ([]LHSFilter, see below), `Headers` (map), `Accept` (*string), `Body` (any, JSON-marshalled).

## LHS filters

The Cycloid API supports LHS bracket filters on `List` routes: `attribute[condition]=value`. The condition is typically `eq`, `rlike`, `gt`, `lt`, etc.

**Rule: all new `List` middleware methods must accept `filters ...LHSFilter` as their last parameter.**

`LHSFilter` is defined in `cmd/cycloid/middleware/lhs_filter.go`:

```go
type LHSFilter struct {
    Attribute string
    Condition string
    Value     string
}
```

Pass filters via the `LHSFilters` field of `Request`. Brackets are kept literal (not percent-encoded) so the API receives `name[eq]=my-project`, not `name%5Beq%5D=my-project`. Regex metacharacters in values (`?`, `*`, `+`, etc.) are also preserved.

```go
// Example: list projects filtered by name prefix
func (m *middleware) ListProjects(org string, filters ...LHSFilter) ([]*models.Project, *http.Response, error) {
    var result []*models.Project
    resp, err := m.GenericRequest(Request{
        Method:       "GET",
        Organization: &org,
        Route:        []string{"organizations", org, "projects"},
        LHSFilters:   filters,
    }, &result)
    ...
}

// Caller usage:
projects, _, err := m.ListProjects(org, middleware.LHSFilter{
    Attribute: "name",
    Condition: "rlike",
    Value:     "proj.*",
})
```

Offline (no-backend) unit tests for LHS filter encoding live in `cmd/cycloid/middleware/offline/lhs_filter_test.go`.

### Return type conventions

| Verb | Return |
|------|--------|
| Get / Create | `(*models.X, *http.Response, error)` |
| List | `([]*models.X, *http.Response, error)` |
| Delete / void | `(*http.Response, error)` |

Always return the `*http.Response` so callers can inspect status codes. Assign `_` if unused.

## Command pattern

```go
var projectTableOptions = printer.Options{
    Columns:    []string{"Canonical", "Name", "Description", "Owner.Username"},
    Identifier: "Canonical",
}

func getProject(cmd *cobra.Command, args []string) error {
    // Step 1: ALL flags first
    org, err := cyargs.GetOrg(cmd)
    if err != nil { return err }
    project, err := cyargs.GetProject(cmd)
    if err != nil { return err }

    // Step 2: API + middleware
    api := common.NewAPI()
    m := middleware.NewMiddleware(api)

    // Step 3: call + print (cyout handles printer selection and error routing)
    result, _, err := m.GetProject(org, project)
    return cyout.PrintWithOptions(cmd, result, err, "unable to get project", projectTableOptions)
}
```

`cyout.PrintWithOptions` handles everything: reads `--output`, picks the printer, routes errors to stderr and results to stdout. For commands with no column customisation, use `cyout.Print(cmd, obj, err, errMsg)`.

## `cyargs` flags

All shared flag definitions live in `internal/cyargs/`. Pattern:

```go
func AddWidgetFlag(cmd *cobra.Command) {
    cmd.Flags().String("widget", "", "Widget canonical")
    _ = cmd.RegisterFlagCompletionFunc("widget", widgetCompletion)
}

func GetWidget(cmd *cobra.Command) (string, error) {
    return cmd.Flags().GetString("widget")
}
```

Register in the command constructor (`NewGetX()`), never inside `RunE`.

## Testing

- **Tests ship with the feature** (see Hard Rule 9): middleware and command changes should include tests in the same PR; extend existing `e2e/*_test.go` files when the resource is already covered there.
- **Middleware tests** (`cmd/cycloid/middleware/*_test.go`): `TestMain` calls `testcfg.NewConfig("middleware")`. Fixtures: `config.Project`, `config.Environment`, `config.Component`, `config.ConfigRepo`, `config.CatalogRepo`.
- **E2E tests** (`e2e/*_test.go`): `executeCommand(args)` runs the real CLI. Same `testcfg` setup.
- E2E tests are **not parallel** (backend git writes are not concurrent-safe).
- `CY_TEST_VERBOSITY=debug` → full HTTP request/response logs. Auth header redacted to `Bearer ***XXXXX` (last 5 chars).

## Config & auth

`common.NewAPI()` resolves config from flags → env vars → config file. Token priority: `--api-key` → `CY_API_KEY` → `CY_API_TOKEN` → per-org token in config file.

## Deeper docs

- `@docs/architecture.md` — HTTP layer, Request struct, auth flow, error taxonomy
- `@docs/pipeline-build-watch-output.md` — `pipeline build trigger --watch` human SSE output; how to disable or remove
- `@docs/adding-a-command.md` — full walkthrough with working example
- `@docs/testing.md` — middleware + e2e test patterns, testcfg deep-dive
- `@docs/middleware-refactor.md` — what changed and why, migration reference

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **cycloid-cli** (11392 symbols, 69497 relationships, 300 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## When GitNexus is Available

- Before modifying a symbol, prefer running `gitnexus_impact({target: "symbolName", direction: "upstream"})` to understand callers and blast radius.
- Before committing, consider running `gitnexus_detect_changes()` to verify the affected scope matches expectations.
- If impact analysis returns HIGH or CRITICAL risk, surface it to the user before proceeding.
- When exploring unfamiliar code, prefer `gitnexus_query({query: "concept"})` over grepping — it returns process-grouped results ranked by relevance.
- For full context on a specific symbol — callers, callees, execution flows — use `gitnexus_context({name: "symbolName"})`.
- Prefer `gitnexus_rename` over find-and-replace for symbol renames — it understands the call graph.

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/cycloid-cli/context` | Codebase overview, check index freshness |
| `gitnexus://repo/cycloid-cli/clusters` | All functional areas |
| `gitnexus://repo/cycloid-cli/processes` | All execution flows |
| `gitnexus://repo/cycloid-cli/process/{name}` | Step-by-step execution trace |

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
