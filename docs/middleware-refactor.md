# Middleware Refactor

## What changed

The middleware layer was rewritten to remove the go-swagger generated operations package (`client/client/`) in favour of a single generic HTTP client (`GenericRequest` in `cmd/cycloid/middleware/generic_client.go`).

### Before

Each API endpoint was represented by a generated operation struct in `client/client/`:

```go
// OLD — do not use
params := operations.NewGetOrganizationsOrgProjectsProjectParams()
params.SetOrganizationCanonical(org)
params.SetProjectCanonical(project)
resp, err := m.api.Client.Organizations.GetOrganizationsOrgProjectsProject(params, auth)
```

### After

All API calls go through `GenericRequest`:

```go
// NEW — current pattern
var result *models.Project
resp, err := m.GenericRequest(Request{
    Method:       "GET",
    Organization: &org,
    Route:        []string{"organizations", org, "projects", project},
}, &result)
```

## Why the change

| Pain point | Old (generated operations) | New (GenericRequest) |
|------------|--------------------------|---------------------|
| Error handling | Inconsistent per-operation types | Single `*APIResponseError` |
| Auth | Per-call auth parameter | Automatic via `Organization` field |
| API changes | Full codegen cycle required | Update route string + model |
| Routing | Opaque generated code | Explicit `Route` slice |
| Testing | Required generated mocks | Direct HTTP testing |
| `{"data":...}` envelope | Handled inconsistently | Always unwrapped in one place |

## New invariants

Every middleware method follows the same contract:

```go
// Delete / void
func (m *middleware) DeleteX(org, x string) (*http.Response, error)

// Get / Create
func (m *middleware) GetX(org, x string) (*models.X, *http.Response, error)

// List
func (m *middleware) ListX(org string) ([]*models.X, *http.Response, error)
```

- Always return `(*http.Response, error)` at minimum so callers can inspect status codes
- Callers assign `_` to the response if they don't need it
- On error, return `nil, resp, err` (include resp so caller can check status)
- Do not validate payloads in middleware — validate in commands or leave it to the API

## The `client/models/` exception

Only the **operations** layer was removed. The **models** layer (`client/models/`) is still generated from `swagger.yaml` and is still used everywhere.

```
client/
  models/   ← STILL generated, DO NOT edit manually
  client/   ← REMOVED (operations layer, do not use)
```

If you see an import like `"github.com/cycloidio/cycloid-cli/client/client/..."`, remove it and replace with `GenericRequest`.

## Local model types

Some models were removed from the swagger spec but are still needed by the CLI. These live directly in the middleware package:

```go
// cmd/cycloid/middleware/http_client.go
type StackVersion struct { ... }
type StackUseCase struct { ... }
```

This is intentional — do not move them to `client/models/`.

## Debug logging

`GenericRequest` emits HTTP debug logs when `CY_VERBOSITY=debug` (set via `--verbosity debug` or the env var). Logs include:

- Method, full URL, headers, body
- Response status, headers, body, elapsed time
- `Authorization` header is redacted: `Bearer ***XXXXX` (last 5 chars only)

During tests: `CY_TEST_VERBOSITY=debug` — `TestMain` propagates it to `CY_VERBOSITY`.

> ⚠️ Debug logs may contain sensitive data in request/response bodies beyond credentials.

## Migration reference

If you find code still using the old generated operations client, replace it:

```go
// OLD — remove
import "github.com/cycloidio/cycloid-cli/client/client/operations"
params := operations.NewGetSomethingParams().WithOrg(org)
resp, err := m.api.Client.Operations.GetSomething(params, auth)
return resp.Payload.Data, nil, err

// NEW — replace with
var result *models.Something
httpResp, err := m.GenericRequest(Request{
    Method:       "GET",
    Organization: &org,
    Route:        []string{"organizations", org, "somethings", id},
}, &result)
return result, httpResp, err
```
