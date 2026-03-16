# Pipeline build watch: human-formatted SSE output

This document describes the **optional presentation layer** for `cy pipeline build trigger --watch` (aliases: `pp build create|run --watch`). It lives in `internal/buildwatch`; `cmd/cycloid/pipelines` only wires flags and calls `buildwatch.Watch`. The layer exists so interactive use is readable; it is **not** required for correctness of the watch loop or API.

If this layer causes regressions (missing lines, wrong parsing, terminal quirks, CI noise), use **Quick disable** below first, then **Full removal** if you want it gone entirely.

---

## User-facing behavior

| Situation | Stream written to stdout |
|-----------|---------------------------|
| `--watch` and **not** `--output json` | **Human**: one-line prefixes (`[build]`, `[task]`, `[stdout]` / `[stderr]`, `[res]` in verbose mode). Concourse JSON events are parsed and filtered. |
| `--watch` and `--output json` | **Raw**: NDJSON lines exactly as in the SSE `data:` payloads (legacy scripting behavior). |
| TTY + human | ANSI colors from `buildwatch.DefaultStreamTheme` in `internal/buildwatch/theme.go`. |
| Non-TTY or piped stdout + human | No ANSI; ANSI in log payloads is stripped. |
| Global `--verbosity debug` (or `-v debug`) + human + `--watch` | Emits extra “plumbing” events (image checks, workers, etc.). No effect when output is raw JSON (`--output json`). |

**Log noise control:** Whitespace-only log payloads and blank lines inside a payload are skipped so empty Concourse log events do not print spacer lines.

**Build link (stderr, before the stream):** With `--watch`, the CLI prints the build id and a console deep link on **stderr** so stdout can stay clean (e.g. raw JSON). The link uses:

`{console_url}/organizations/{org}/projects/{project}/environments/{env}/components/{component}/pipelines/{pipeline}/jobs/{job}/builds/{build_id}`

The default `console_url` is `https://console.cycloid.io`. Override with **`CY_CONSOLE_URL`** only (no CLI flag). If `CY_CONSOLE_URL` is set to an empty value, a stderr line explains that no link can be printed.

**Ctrl+C reminder:** Before the stream, stderr also notes: first Ctrl+C requests a remote build abort; second exits watch immediately (see below).

## Ctrl+C while watching

Handled in `internal/buildwatch/watch.go` (not the cobra layer):

1. **First Ctrl+C:** sends a **build abort** request (`AbortBuild`) so the remote build can stop gracefully. A short hint is printed to stderr (via `StatusWriter`): press Ctrl+C again to leave watch immediately.
2. **Second Ctrl+C:** **cancels** the local watch (stops polling and the event stream). The process exits with code **130** (`watch interrupted`), matching the usual `128 + SIGINT` convention.

Disable this for tests with `Options.DisableInterruptHandler`.

---

## Code map (files to inspect or delete)

| File | Role |
|------|------|
| `cmd/cycloid/pipelines/build_trigger.go` | Cobra only: chooses `buildwatch.OutputHuman` vs `OutputRaw`, TTY theme, `StripLogANSI`, verbosity→`Verbose`, calls `buildwatch.Watch`. |
| `internal/buildwatch/watch.go` | `Watch`, `Client`, `Options`, `ExitError`, SSE consumption, reconnect loop, polling, Ctrl+C handling. |
| `internal/buildwatch/format.go` | Human line rendering, `formatHumanBuildEvent`, `paintLogLines`, ANSI stripping in payloads. |
| `internal/buildwatch/theme.go` | `StreamTheme`, `DefaultStreamTheme` (edit here to change colors only). |
| `internal/buildwatch/events.go` | JSON structs for Concourse event envelopes. |
| `internal/buildwatch/watch_test.go` | Tests for formatters and watch wiring. |

Middleware `OpenBuildEventsStream` (`cmd/cycloid/middleware/component_pipelines_jobs_builds.go`) is unchanged: it still returns `text/event-stream`.

---

## Quick disable (keep package, force raw stream)

To make **`--watch` always behave like raw NDJSON** regardless of `--output` (fastest rollback):

1. Open `cmd/cycloid/pipelines/build_trigger.go`.
2. In the `if watch {` block, force raw output mode, for example:

```go
outMode := buildwatch.OutputRaw
// Was: human unless --output json; temporarily forcing raw — see docs/pipeline-build-watch-output.md
```

3. You can drop the `theme` / `stripLogANSI` / `IsTerminalWriter` block when `outMode` is always `OutputRaw` (optional cleanup).

4. Run `make format && make lint` and `go test ./internal/buildwatch/... ./cmd/cycloid/pipelines/...`.

Debug verbosity then only affects HTTP logging elsewhere, not the watch stream (already raw).

---

## Full removal (restore passthrough-only watch)

1. **Delete** (or revert in git) the package directory:
   - `internal/buildwatch/` (all `.go` files)

2. **`cmd/cycloid/pipelines/build_trigger.go`**
   - Remove `internal/buildwatch` import.
   - Inline a minimal watch loop **or** copy a slim `watch.go` back under `cmd/cycloid/pipelines` that only passthrough-writes SSE `data:` payloads (see git history before `internal/buildwatch`).

3. **`go.mod`**
   - Run `go mod tidy` if `golang.org/x/term` is no longer referenced.

4. **Docs**
   - Delete this file and remove links from `docs/architecture.md`, `CLAUDE.md`, and `AGENTS.md`.

---

## Flags reference

| Flag | Meaning |
|------|---------|
| `--watch` | Wait for build completion while streaming build events. |
| `--timeout` | With `--watch`, max seconds before giving up (see `--watch-cancel-on-timeout`). |
| `--watch-cancel-on-timeout` | With `--watch`, call abort when `--timeout` elapses. |
| `--output json` | Watch stream stays machine-readable NDJSON. |
| `--verbosity debug` | With human watch output, print all parsed event types. |

---

## Maintenance notes

- **Do not** tie this to `cyargs` project `--color`; watch colors are isolated in `internal/buildwatch/theme.go`.
- Parser failures on a `data:` line fall back to writing that line unchanged (plus newline), so unknown event shapes still surface as text.
- Changing Concourse event schemas may require updates to `internal/buildwatch/events.go` and the `switch` in `formatHumanBuildEvent` (`format.go`).
