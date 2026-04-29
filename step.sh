#!/usr/bin/env bash

set -eo pipefail
set -x

# download the Bitrise Build Cache CLI
export BITRISE_BUILD_CACHE_CLI_VERSION="v2.4.1"
curl --retry 5 -m 30 -sSfL 'https://raw.githubusercontent.com/bitrise-io/bitrise-build-cache-cli/main/install/installer.sh' | sh -s -- -b /tmp/bin -d $BITRISE_BUILD_CACHE_CLI_VERSION || true

# Fall back to Artifact Registry if the download failed
if [ ! -f /tmp/bin/bitrise-build-cache ]; then
  echo "Failed to download Bitrise Build Cache CLI, trying Artifact Registry ..."

  version="${BITRISE_BUILD_CACHE_CLI_VERSION#v}"
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  arch=$(uname -m | sed 's/x86_64/amd64/' | sed 's/aarch64/arm64/')
  package="bitrise-build-cache_${os}_${arch}.tar.gz"
  filename="bitrise-build-cache_${version}_${os}_${arch}.tar.gz"

  filepath="$package:$version:$filename"

  echo "Downloading Bitrise Build Cache CLI from Artifact Registry: ${filepath}"

  curl --retry 5 -m 60 -sSfL "https://artifactregistry.googleapis.com/download/v1/projects/ip-build-cache-prod/locations/us-central1/repositories/build-cache-cli-releases/files/${filepath}:download?alt=media" -o $package
  tar -xzf "$package"
  mkdir -p /tmp/bin
  mv "bitrise-build-cache" /tmp/bin/bitrise-build-cache
  rm -rf "$package"
fi

if [ ! -f /tmp/bin/bitrise-build-cache ]; then
  echo "Failed to download Bitrise Build Cache CLI, exiting."
  exit 1
fi

# Build the activate gradle-mirrors invocation, passing through the flags only when explicitly set.
args=("activate" "gradle-mirrors" "--debug=$verbose")

if [ "$mavencentral" = "true" ]; then
  args+=("--mavencentral")
fi

if [ "$mavencentral_apache" = "true" ]; then
  args+=("--mavencentral-apache")
fi

if [ "$google" = "true" ]; then
  args+=("--google")
fi

/tmp/bin/bitrise-build-cache "${args[@]}"
