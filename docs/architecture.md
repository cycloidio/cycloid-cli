# Architecture

## Request lifecycle

```
cmd/cycloid/<feature>/verb.go
        │  (cobra RunE)
        ▼
cyargs.Get*()              ← parse ALL flags first (required — see CLAUDE.md Hard Rules)
        │
common.NewAPI()            ← build API config (URL, token) from flags/env/config file
        │
middleware.NewMiddleware()  ← construct middleware struct with HTTP client
        │
m.GetX / m.ListX / ...    ← middleware method in cmd/cycloid/middleware/<feature>.go
        │
m.GenericRequest(Request{...}, &result)
        │
        ▼
HTTP → Cycloid REST API
        │
        ▼
JSON {"data": <payload>}   ← GenericRequest unwraps envelope; &result receives payload
        │
printer.SmartPrint()       ← stdout (success) or stderr (error)
```

## The `Request` struct

Defined in `cmd/cycloid/middleware/http_client.go`:

```go
type Request struct {
    Method       string
    Organization *string           // used for auth token lookup; nil = no org context
    NoAuth       bool              // set true to skip Authorization header
    Route        []string          // joined onto base URL path: ["organizations", org, "projects"]
    Query        any               // struct with `url` tags, or url.Values
    Headers      map[string]string // extra headers merged into request
    Accept       *string           // overrides default Accept header
    Body         any               // JSON-marshalled when non-nil
}
```

`Route` segments are path-joined onto the base URL (e.g., `CY_API_URL`). The route should not start with `/`.

## `GenericRequest` behaviour

Defined in `cmd/cycloid/middleware/generic_client.go`.

1. Builds the full URL from `m.api.Config.URL` + `req.Route`
2. Encodes `req.Query` via struct tags (`url:"param_name"`)
3. JSON-marshals `req.Body`
4. Sets `Content-Type: application/json` and (unless `NoAuth`) `Authorization: Bearer <token>`
5. Executes the HTTP call
6. On non-2xx: returns `*APIResponseError`
7. On 2xx: unwraps the `{"data": ...}` JSON envelope into `response` (the second argument)
   — callers pass `&result` directly, not a `struct{ Data *X }` wrapper

```go
// Correct pattern:
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

## Envelope unwrapping

The Cycloid API wraps all responses:

```json
{ "data": { ... } }
```

`GenericRequest` strips this envelope before deserializing into `response`. If the response body is not envelope-shaped, it falls back to direct unmarshal.

Pass `nil` as `response` to discard the body (e.g., DELETE calls).

## Authentication

`m.api.GetToken(org)` resolves the bearer token in this priority order:

1. `--api-key` flag
2. `CY_API_KEY` env var
3. `CY_API_TOKEN` env var (legacy)
4. Per-org token stored in the config file (`~/.cy/config.yml`)

Pass `Organization: &org` in `Request` to allow token lookup. Set `NoAuth: true` for unauthenticated endpoints (e.g., login).

## Error taxonomy

| Type | Cause | Go type |
|------|-------|---------|
| API error | Server returned non-2xx | `*APIResponseError` |
| Network error | Transport failure (DNS, TLS, timeout) | `*url.Error` or stdlib |
| Unexpected error | Anything else | `error` (do not wrap) |

### `APIResponseError`

```go
type APIResponseError struct {
    StatusCode int
    Status     string
    Body       []byte                 // raw response body
    Payload    *models.ErrorPayload   // parsed if body was valid JSON error
    Path       string                 // request path (+ query) for fallback errors
}

// Error() format:
// - payload message available: "API error 422: <message from payload>"
// - fallback body/path:       "API error 422 on "/path?query": <raw body>"
```

Check with `errors.As`:

```go
var apiErr *middleware.APIResponseError
if errors.As(err, &apiErr) {
    if apiErr.StatusCode == 409 {
        // conflict
    }
}
```

Common status codes: 400 bad request, 401 unauthorized, 403 forbidden, 404 not found, 409 conflict, 422 unprocessable entity.

### JSON output and API error diagnostics

Errors that implement `printer.ErrHTTPResponse` attach the HTTP status code and raw response body (`*APIResponseError` for non-2xx, plus the decode error type returned by `GenericRequest` when a 2xx body cannot be unmarshaled). When `--output json` is used, the JSON printer first tries to marshal the value normally. If marshaling fails but the value is an error satisfying `ErrHTTPResponse`, it prints a small JSON object instead: `cli_marshal_error`, `http_status`, `api_response_preview` (first 10 lines of the body), and optionally `request_path` when the error also implements `printer.RequestPather` (as `*APIResponseError` does).

## Why `client/client/` is unused

The repository used to use the go-swagger generated operations package (`client/client/`). It was removed in the middleware refactor (see `docs/middleware-refactor.md`) in favour of `GenericRequest`:

- Generated operations had inconsistent error handling
- Every API change required re-running the full swagger codegen cycle
- `GenericRequest` gives explicit control over routing, auth, headers, and response decoding

The `client/models/` package (data types) is still auto-generated from `swagger.yaml` and must not be edited manually.

## Pipeline build watch (human SSE formatting)

`cy pipeline build trigger --watch` streams build events while polling until the build finishes. Human vs raw NDJSON formatting lives in `internal/buildwatch`; the cobra command only passes options and calls `buildwatch.Watch`. If it causes problems, see [pipeline-build-watch-output.md](./pipeline-build-watch-output.md) for how to disable or remove it without touching the HTTP client.
