{ pkgs, ... }:

# devenv.sh dev environment for the Cycloid CLI (`cy`).
# One reproducible Go toolchain shared by local dev (`devenv shell` / `direnv
# allow`) and CI (the self-hosted runner's lint-build). Supersedes the old
# `flake.nix` mkShell, which declared `golangci-lint`/`gci` commented-out and
# omitted `goimports`/`shfmt`/`shellcheck` entirely — so `make format lint`
# could not run from inside it. Git hooks are declared below via devenv's
# git-hooks integration (auto-installed on shell entry).
let
  # `make client-generate` post-processes the generated swagger client with a
  # small PyYAML script (see scripts/ + Makefile client-generate target).
  swaggerPython = pkgs.python312.withPackages (p: with p; [ pyyaml ]);
in
{
  packages = with pkgs; [
    # Go toolchain — pinned to match `go 1.25` in go.mod. Plain package (not the
    # `languages.go` module) on purpose: that module also injects gopls/delve,
    # and unstable's gopls now requires go >= 1.26, which would force a toolchain
    # ahead of go.mod. Editors bring their own gopls; the shell stays minimal,
    # exactly like the flake it replaces.
    go_1_25

    gnumake # the Makefile drives build / lint / format / test
    git

    # Swagger client generation (`make client-generate*`).
    go-swagger
    swaggerPython

    # Lint + format — every tool the Makefile's lint-go/lint-sh/format-go/
    # format-sh targets and the format-and-lint pre-commit hook invoke.
    golangci-lint # lint-go: golangci-lint run -v   (.golangci.yml)
    gci # format-go: gci write ...
    gotools # provides goimports (format-go: goimports -w)
    shellcheck # lint-sh
    shfmt # format-sh

    # Dev tooling.
    awscli2
    docker-client # `docker compose` for `make be-start` / e2e; daemon is host-provided
  ];

  # Pre-commit hooks via devenv's git-hooks integration — declared here and
  # auto-installed on shell entry, replacing the hand-written
  # .pre-commit-config.yaml and `make install`'s hook wiring.
  git-hooks.hooks = {
    # The repo's canonical gate: `make format` (gci + goimports + shfmt) then
    # `make lint` (golangci-lint + shellcheck) — Hard Rule 8. The Makefile stays
    # the single source of truth, shared with CI's lint-build job.
    format-and-lint = {
      enable = true;
      name = "format and lint";
      entry = "make format lint";
      pass_filenames = false;
    };
    # Block direct commits to the shared integration branches.
    no-commit-to-branch = {
      enable = true;
      settings.branch = [
        "develop"
        "master"
      ];
    };
    # NOTE: no secret-scanner hook. The old detect-private-key hook was never
    # actually enforced (the repo commits many intentional test/example keys —
    # compose.yml fixtures, e2e SSH keys, the AKIA…EXAMPLE AWS key), and both it
    # and ripsecrets flag all of those. A scanner here needs a fragile allowlist
    # for zero real benefit, so it's intentionally omitted.
  };

  enterShell = ''
    echo "cycloid-cli (cy) dev shell"
    echo "  go:            $(go version | awk '{print $3}')"
    echo "  golangci-lint: $(golangci-lint version --short 2>/dev/null || golangci-lint --version 2>/dev/null | head -1)"
    echo "  go-swagger:    $(swagger version 2>/dev/null | head -1)"
    echo
    echo "  make build            # cross-compile all cy binaries"
    echo "  make format lint      # what the pre-commit hook runs"
    echo "  make test             # go test ./..."
  '';
}
