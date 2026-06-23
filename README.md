# Cycloid CLI (cy)

Command-line interface for the [Cycloid](https://cycloid.io) platform.

## Installation

### Go install (recommended for developers)

```sh
go install github.com/cycloidio/cycloid-cli@latest
```

This installs a lightweight shim that automatically downloads the correct platform binary on first run.

### Pre-built binaries

Download pre-built binaries from the [releases page](https://github.com/cycloidio/cycloid-cli/releases).

### In pipelines

Use the [`cycloid/cycloid-toolkit`](https://hub.docker.com/r/cycloid/cycloid-toolkit) Docker image, which includes a wrapper that automatically downloads the correct CLI version matching your Cycloid API.

## Usage

```sh
# First run downloads the real binary (~12 MB), then executes it
cy organizations list

# Pin a specific version
CY_VERSION=v6.10.24 cy --version

# Check shim version
cy --shim-version
```

## How it works

The `go install` binary is a thin installer shim (~3 MB). On first run it:

1. Determines your platform (linux/darwin/windows, amd64/arm64)
2. Downloads the matching pre-built `cy` binary from GitHub Releases
3. Caches it at `~/.cycloid/bin/cy-<version>`
4. Exec's the real binary with all your arguments

Subsequent runs use the cached binary directly (no network call).

## Documentation

Full CLI documentation: [docs.cycloid.io/reference/cli](https://docs.cycloid.io/reference/cli/)

## License

[MIT](LICENSE)
