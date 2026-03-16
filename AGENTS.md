# AGENTS.md

This file mirrors `CLAUDE.md` for LLM agents that use the AGENTS.md convention (e.g., OpenAI Codex, Gemini Code Assist). Keep in sync with `CLAUDE.md` when updating either file.

---

## Commands

```bash
# Build
make build # build all platform binaries
go build -o cy . # quick local build

# Test (requires local backend)
make be-start # start backend via docker compose
go test ./... # run all tests
make test # same
make be-stop # stop backend

# Run a specific test
go test ./e2e/... -run TestProjects
go test ./cmd/cycloid/middleware/... -run TestGetProject

# Lint & format
make lint # golangci-lint + shellcheck
make format # gci + goimports + shfmt

# Client regeneration (when swagger.yaml changes)
make client-generate # regenerates ./client/ from swagger.yaml

# Changelog
make new-changelog-entry # add unreleased changelog entry (uses changie via docker)
```

## Architecture

```
cmd/cycloid/<feature>/*.go → cmd/cycloid/middleware/ → GenericRequest()
 (cobra commands) (Middleware interface) (generic_client.go)
```

### Key packages

- **`client/models/`** — Auto-generated go-swagger models from `swagger.yaml`. Never edit manually. Regenerate with `make client-generate`.
- **`client/client/`** — Intentionally NOT used. The generated operations package was removed in the middleware refactor. Use `GenericRequest` instead.
- **`cmd/cycloid/middleware/`** — The `Middleware` interface (`middleware.go`) makes HTTP calls via `GenericRequest` (`generic_client.go`). Each feature gets its own file (e.g., `organization_projects.go`).
- **`cmd/cycloid/<feature>/`** — Cobra command definitions. Each feature directory has: `cmd.go` (registers subcommands), plus one file per verb (`list.go`, `get.go`, `create.go`, `update.go`, `delete.go`), and `common.go` for shared logic.
- **`internal/cyargs/`** — All shared flag definitions and completion functions. Every flag used by multiple commands must be declared here.
- **`printer/`** — Output formatting. Use `printer.SmartPrint(p, obj, err, errStr, opts, writer)` — success to `cmd.OutOrStdout()`, errors to `cmd.OutOrStderr()`.
- **`e2e/`** — End-to-end tests that run real CLI commands against a live backend.

## Hard Rules

These are invariants that LLM agents and new contributors must not violate:

1. **NEVER edit `client/models/`** — auto-generated. Run `make client-generate` to update.
2. **NEVER import or use `client/client/`** — deprecated. Use `GenericRequest` in middleware methods.
3. **NEVER call `NewMiddleware` outside a cobra `RunE` function** — not at package init, not in tests directly.
4. **ALWAYS parse ALL flags via `cyargs.Get*` BEFORE calling `common.NewAPI()` and `NewMiddleware()`** — `GenericRequest` reads `verbosity` from Viper at call time; unparsed flags produce stale values.
5. **ALWAYS add shared flags to `internal/cyargs/`** — never inline a flag used by more than one command.
6. **ALWAYS use `printer.SmartPrint`** — errors go to `OutOrStderr()`, results go to `OutOrStdout()`.
7. **E2E tests require a running backend** (`make be-start`). Never run e2e in parallel.
8. **Run `make format && make lint` after every code change.**

## `GenericRequest` pattern

`GenericRequest` handles auth, JSON marshaling, and `{"data":...}` envelope unwrapping. Pass `&result` directly — do not wrap in a `struct{ Data *X }`.

```go
// Standard middleware method pattern:
func (m *middleware) GetProject(org, project string) (*models.Project, *http.Response, error) {
 var result *models.Project
 resp, err := m.GenericRequest(Request{
 Method: "GET",
 Organization: &org,
 Route: []string{"organizations", org, "projects", project},
 }, &result)
 if err != nil {
 return nil, resp, err
 }
 return result, resp, nil
}
```

`Request` fields: `Method`, `Organization` (*string, for auth), `NoAuth` (bool), `Route` ([]string), `Query` (struct with `url` tags), `Headers` (map), `Accept` (*string), `Body` (any, JSON-marshalled).

### Return type conventions

| Verb | Return |
|------|--------|
| Get / Create | `(*models.X, *http.Response, error)` |
| List | `([]*models.X, *http.Response, error)` |
| Delete / void | `(*http.Response, error)` |

Always return the `*http.Response` so callers can inspect status codes. Assign `_` if unused.

## Command pattern

```go
func getProject(cmd *cobra.Command, args []string) error {
 // Step 1: ALL flags first
 org, err := cyargs.GetOrg(cmd)
 if err != nil { return err }
 project, err := cyargs.GetProject(cmd)
 if err != nil { return err }
 output, err := cmd.Flags().GetString("output")
 if err != nil { return errors.Wrap(err, "unable to get output flag") }

 // Step 2: printer
 p, err := factory.GetPrinter(output)
 if err != nil { return errors.Wrap(err, "unable to get printer") }

 // Step 3: API + middleware
 api := common.NewAPI()
 m := middleware.NewMiddleware(api)

 // Step 4: call + print
 result, _, err := m.GetProject(org, project)
 if err != nil {
 return printer.SmartPrint(p, nil, err, "unable to get project", printer.Options{}, cmd.OutOrStderr())
 }
 return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
```

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

- **Middleware tests** (`cmd/cycloid/middleware/*_test.go`): `TestMain` calls `testcfg.NewConfig("middleware")`. Fixtures: `config.Project`, `config.Environment`, `config.Component`, `config.ConfigRepo`, `config.CatalogRepo`.
- **E2E tests** (`e2e/*_test.go`): `executeCommand(args)` runs the real CLI. Same `testcfg` setup.
- E2E tests are **not parallel** (backend git writes are not concurrent-safe).
- `CY_TEST_VERBOSITY=debug` → full HTTP request/response logs. Auth header redacted to `Bearer ***XXXXX` (last 5 chars).

## Config & auth

`common.NewAPI()` resolves config from flags → env vars → config file. Token priority: `--api-key` → `CY_API_KEY` → `CY_API_TOKEN` → per-org token in config file.

## Deeper docs

- `@docs/architecture.md` — HTTP layer, Request struct, auth flow, error taxonomy
- `@docs/adding-a-command.md` — full walkthrough with working example
- `@docs/testing.md` — middleware + e2e test patterns, testcfg deep-dive
- `@docs/middleware-refactor.md` — what changed and why, migration reference
