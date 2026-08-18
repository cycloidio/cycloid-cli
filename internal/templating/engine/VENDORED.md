# Vendored interpolation engine — TEMPORARY

This package is a **thin-adapted copy** of the Cycloid interpolation engine from
`youdeploy-http-api`, pulled in so the CLI can render templates **offline** with
no backend. It exists only until the CLI→backend merge.

## Source

- Repo: `cycloidio/youdeploy-http-api`
- Commit: `39d97e36b1dc683fe3582416888fb26df4f0da70`
- Files:
  - `utils/interpolator.go` → `interpolator.go`
  - `utils/interpolator_entity_string.go` → `interpolator_entity_string.go` (verbatim)
  - `utils/helmutils/func_map.go` → `func_map.go` (verbatim, repackaged)
  - `utils/interpolator_error.go` → `interpolator_error.go` (adapted)
  - `services/youdeploy/svccat/version`→ `version.go` (minimal local reimplementation)

## Adaptations

The only changes from upstream are error plumbing: the backend's
`yderr`/`errtmpl` taxonomy (~4.8k lines, DB/service-coupled) is replaced with
stdlib `errors`/`fmt`. **Rendered output is identical** — only error *types* and
*wording* differ. This is what the render-parity test guards (output, not error
internals).

## On the CLI→backend merge

Delete this whole directory and import the engine directly from the backend
(`utils.Interpolator`). Re-point `internal/templating` at it. The parity test
becomes redundant for the engine half at that point.

## Do not

- Add features here. Behavioural changes belong upstream, then re-vendor.
- Re-introduce `yderr`/`errtmpl` — keep the adapter surface minimal.
