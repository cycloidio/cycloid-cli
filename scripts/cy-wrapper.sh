#!/usr/bin/env bash

# Warning: this script is expected to be used in cycloid-toolkit docker image.
# We do not recommand to use it in another context.

if test -z "$BASH_VERSION"; then
  echo "Please run this script using bash, not sh or any other shell." >&2
  exit 1
fi

export CY_API_URL="${CY_API_URL:-https://http-api.cycloid.io}"
export CY_BINARIES_PATH="${CY_BINARIES_PATH:-/usr/local/bin}"

# Ensure we have Cycloid directory created
if ! [ -d "$CY_BINARIES_PATH" ]; then
  mkdir -p "$CY_BINARIES_PATH"
fi

# If specified in the commandline use it
if [[ $(echo "$@" | grep -E --  "--api-url(=|\s+)([^\]+)") ]]; then
  CY_API_URL=$(echo "$@" | sed -r "s/.*--api-url(=|\s+)([^ ]+).*/\2/")
fi

# Remove trailing /
CY_API_URL=${CY_API_URL%/}

# Get Cycloid API version
export CY_VERSION=$(curl -s "${CY_API_URL}/version" | jq -r .data.version)

if [[ -z "$CY_VERSION" ]]; then
  echo "Error: Unable to get Cycloid API version on ${CY_API_URL}/version" >&2
  exit 1
fi

# Download the binary if not present
export CY_BINARY="${CY_BINARY:-"${CY_BINARIES_PATH}/cy-${CY_VERSION}"}"
if ! [[ -f "${CY_BINARY}" ]]; then

  # Download the exact CLI version
  CY_URL="https://github.com/cycloidio/cycloid-cli/releases/download/v${CY_VERSION}/cy"
  wget -q -O "${CY_BINARY}" "$CY_URL"
  STATUS=$?

  if [ $STATUS != 0 ]; then
    rm -f "${CY_BINARY}"
  fi

  # In case of 404, download RC CLI version
  if [ $STATUS == 8 ]; then
    echo "Warning: Unable to download CLI version ${CY_VERSION}. Fallback to RC version" >&2
    CY_BINARY="${CY_BINARIES_PATH}/cy-${CY_VERSION}-rc"
    CY_URL="https://github.com/cycloidio/cycloid-cli/releases/download/v${CY_VERSION}-rc/cy"
    wget -q -O "${CY_BINARY}" "$CY_URL"
    STATUS=$?
    if [ $STATUS != 0 ]; then
      rm -f "${CY_BINARY}"
    fi
  fi

  # In case of 404, fallback on latest develop version
  if [ $STATUS == 8 ]; then
    echo "Warning: Unable to download CLI version ${CY_VERSION}-rc. Fallback to latest develop version" >&2
    CY_BINARY="${CY_BINARIES_PATH}/cy-latest"
    CY_URL="https://github.com/cycloidio/cycloid-cli/releases/download/v0.0-dev/cy"
    wget -q -O "${CY_BINARY}" "$CY_URL"
    STATUS=$?
    if [ $STATUS != 0 ]; then
      rm -f "${CY_BINARY}"
    fi
  fi

  if [ $STATUS != 0 ]; then
    echo "Error: Unable to download Cycloid CLI from github ${CY_URL}" >&2
    exit 1
  fi
  chmod +x "${CY_BINARY}"
fi

# Run Cycloid CLI
exec "${CY_BINARY}" "$@"
