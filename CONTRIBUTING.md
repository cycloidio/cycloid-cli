# Contributing

## Adding a feature or command

See [`docs/adding-a-command.md`](docs/adding-a-command.md) for a full walkthrough including middleware methods, cobra commands, flags, and tests.

## Running tests

See [`docs/testing.md`](docs/testing.md) for setup, environment variables, and test patterns.

Quick start:

```bash
make be-start          # start the local backend
go test ./...          # run all tests
make be-stop
```

## Code style

```bash
make format            # gci + goimports + shfmt
make lint              # golangci-lint + shellcheck
```

Always run both before submitting a PR.

## Changelog

Each PR should include a changelog entry:

```bash
make new-changelog-entry
```

This runs `changie` via docker and creates a file in `changelog/unreleased/`.

## Architecture overview

See [`docs/architecture.md`](docs/architecture.md) and [`CLAUDE.md`](CLAUDE.md).
