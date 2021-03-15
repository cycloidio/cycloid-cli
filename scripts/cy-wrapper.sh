#!/usr/bin/env bash

# Warning: this script is expected to be used in cycloid-toolkit docker image.
# We do not recommand to use it in another context.

if test -z "$BASH_VERSION"; then
  echo "Please run this script using bash, not sh or any other shell." >&2
  exit 1
fi

export CY_API_URL="${CY_API_URL:-https://http-api.cycloid.io}"
export CY_BINARIES_PATH="${CY_BINARIES_PATH:-/usr/local/bin}"


# Compating version, this is used when there is no CLI matching your API version.
# We compare your version and the one released to find the closest n-1 version
vercomp () {
    # Src https://stackoverflow.com/questions/4023830/how-to-compare-two-strings-in-dot-separated-version-format-in-bash
    # Return code
    # 0) =
    # 1) >
    # 2) <
    if [[ $1 == $2 ]]
    then
        return 0
    fi
    local IFS=.
    local i ver1=($1) ver2=($2)
    # fill empty fields in ver1 with zeros
    for ((i=${#ver1[@]}; i<${#ver2[@]}; i++))
    do
        ver1[i]=0
    done
    for ((i=0; i<${#ver1[@]}; i++))
    do
        if [[ -z ${ver2[i]} ]]
        then
            # fill empty fields in ver2 with zeros
            ver2[i]=0
        fi
        if ((10#${ver1[i]} > 10#${ver2[i]}))
        then
            return 1
        fi
        if ((10#${ver1[i]} < 10#${ver2[i]}))
        then
            return 2
        fi
    done
    return 0
}

# Remove extra prefix and suffix from version to compare them
format_version () {
  echo $1 | sed 's/^v//;s/-rc.*$//'
  return
}

# Find the closest n-1 version. This function is used only after we already check if version and version-rc are not available.
# So we only expect to find and take the n-1 version
find_version_below () {
  api_version=$(format_version $1)
  for cli_release in $(curl --silent "https://api.github.com/repos/cycloidio/cycloid-cli/releases" | jq -r '.[] | .name'); do
    cli_version=$(format_version $cli_release)
    # Ignoring the dev version
    if [[ "$cli_version" == "0.0-dev" ]]; then
      continue
    fi

    # Take the first CLI version lower than the API version
    # ret 1 means >
    vercomp $api_version $cli_version
    if [ $? -eq 1 ]; then
      echo $cli_release
      return 0
    fi
  done
  return 2
}

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
    CY_LOWER_VERSION=$(find_version_below ${CY_VERSION})
    echo "Warning: Unable to download CLI version ${CY_VERSION}-rc. Fallback to the closest n-1 version ${CY_LOWER_VERSION}" >&2
    # Removing the v prefix as we don't let it in the binary name
    CY_BINARY="${CY_BINARIES_PATH}/cy-$(echo $CY_LOWER_VERSION | sed 's/^v//')"
    CY_URL="https://github.com/cycloidio/cycloid-cli/releases/download/${CY_LOWER_VERSION}/cy"
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
