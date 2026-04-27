#!/usr/bin/env bash
# generate-env.sh — Resolve .env from .env.sample using the SaaS API.
#
# Prerequisites:
#   export CY_SAAS_API_KEY=<your Cycloid SaaS API key>
#
# Usage:
#   ./scripts/generate-env.sh
#
set -euo pipefail

if [[ -z "${CY_SAAS_API_KEY:-}" ]]; then
    echo "ERROR: CY_SAAS_API_KEY is not set."
    echo "Export it first: export CY_SAAS_API_KEY=<your-key>"
    exit 1
fi

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$REPO_ROOT"

# Build the CLI if needed
if [[ ! -x ./cy ]] || [[ main.go -nt ./cy ]]; then
    echo "==> Building cy..."
    go build -o cy .
fi

echo "==> Resolving .env from .env.sample..."
CY_API_KEY="$CY_SAAS_API_KEY" ./cy uri interpolate .env.sample > .env

echo "==> .env generated. Use: source .env"
