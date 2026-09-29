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

## Known drift from the source commit

The recorded commit is the last **full** re-vendor. Two things have happened
since, neither of them a re-vendor:

- `current_user_email` (PROD-881) was added to `interpolator.go` here and
  upstream **in the same commit**, so the enum, the struct field, `dataMap()`
  and the generated `interpolator_entity_string.go` are in sync. Adding an enum
  entry also means bumping the upper bound of the `KnownKeys()` loop in
  `known_keys.go` — it iterates `org..<last entity>` and a missed bump makes the
  CLI report the new variable as an unknown one.
- `env_vars`/`env_providers` are **not** in sync, and this one changes
  rendering, not just byte-identity. Upstream (PROD-647) installs them
  unconditionally, empty when unset; this copy still guards them behind
  `len() > 0`, and the CLI always builds the interpolator with nil EnvVars.
  `KnownKeys()` lists the four names, so a *bare* `($ .env_vars $)` still gets
  a placeholder — but the dotted `($ .env_vars.<key> $)` form, which is the
  only one the docs actually document, gets none: `templating.Render`'s
  `reBareRef` matches single-segment refs only. Under the engine's
  `missingkey=zero` the absent key then resolves to a zero `interface{}` and
  the `.<key>` lookup on it **errors the whole render**, where the backend
  renders `<no value>`. Re-vendoring the PROD-647 change closes it.
  `templating_test.go` only exercises `.env_vars.region` *with* env_vars
  supplied, so the unset case is uncovered.

## On the CLI→backend merge

Delete this whole directory and import the engine directly from the backend
(`utils.Interpolator`). Re-point `internal/templating` at it. The parity test
becomes redundant for the engine half at that point.

## Do not

- Add features here. Behavioural changes belong upstream, then re-vendor.
- Re-introduce `yderr`/`errtmpl` — keep the adapter surface minimal.
