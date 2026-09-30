#!/bin/bash
# Updates every file that pins the Go toolchain used by CI and the snippets validators.
# Invoked by .github/workflows/check-go-versions.yml as `update-go-versions.sh <latest> <penultimate>`.

set -euo pipefail

LATEST_VERSION=${1:-}
PENULTIMATE_VERSION=${2:-}

if [ -z "${LATEST_VERSION}" ] || [ -z "${PENULTIMATE_VERSION}" ]; then
  echo "Usage: $0 <latest Go version> <penultimate Go version>"
  exit 1
fi

cd "$(dirname "$0")/.."

versions_file=.github/variables/go-versions.env
toolchains_file=snippets/validators/shared/versions/toolchains.env

sed -i \
  -e "s#^latest=.*#latest=${LATEST_VERSION}#" \
  -e "s#^penultimate=.*#penultimate=${PENULTIMATE_VERSION}#" \
  "${versions_file}"
sed -i -e "s#^GO_VERSION=.*#GO_VERSION=${LATEST_VERSION}#" "${toolchains_file}"

grep -qxF "latest=${LATEST_VERSION}" "${versions_file}"
grep -qxF "penultimate=${PENULTIMATE_VERSION}" "${versions_file}"
grep -qxF "GO_VERSION=${LATEST_VERSION}" "${toolchains_file}"

echo "updated ${versions_file} and ${toolchains_file}"
