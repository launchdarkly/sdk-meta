#!/bin/bash

# Wires the genspecs pipeline together. Mirrors scripts/generate-products.sh.
#
# Steps:
#   1. sync-repos: ensure every repo from products/repos.json (plus sdk-specs
#      and sdk-test-harness) is present locally and fast-forwarded if safe.
#   2. catalog: parse sdk-specs into products/specs.json.
#   3. harness: parse sdk-test-harness + each SDK's testharness-suppressions
#      into products/harness_signals.json.
#   4. judge: classify every (sdk, spec) cell into products/spec_support.json.
#      Requires ANTHROPIC_API_KEY when --provider=anthropic (the default when
#      that env var is set). Set GENSPECS_PROVIDER=noop to skip the LLM step
#      and produce an unknown-only matrix that's still useful for verifying
#      the rest of the pipeline.

set -euo pipefail

ROOT="${SDK_REPOS_ROOT:-${HOME}/code/launchdarkly}"
PROVIDER="${GENSPECS_PROVIDER:-}"
EXTRA_ARGS=()
while [[ $# -gt 0 ]]; do
  case "$1" in
    --provider) PROVIDER="$2"; shift 2 ;;
    --skip-sync) SKIP_SYNC=1; shift ;;
    --skip-judge) SKIP_JUDGE=1; shift ;;
    *) EXTRA_ARGS+=("$1"); shift ;;
  esac
done

cd "$(dirname "$0")/.."

if [[ -z "${SKIP_SYNC:-}" ]]; then
  echo "==> sync-repos"
  go -C tool run ./cmd/genspecs sync-repos --sdk-repos-root "$ROOT"
fi

echo "==> catalog"
go -C tool run ./cmd/genspecs catalog \
  --specs-repo "$ROOT/sdk-specs" \
  --out products/specs.json

echo "==> harness"
go -C tool run ./cmd/genspecs harness \
  --harness-repo "$ROOT/sdk-test-harness" \
  --sdk-repos-root "$ROOT" \
  --repos-json products/repos.json \
  --out products/harness_signals.json

if [[ -z "${SKIP_JUDGE:-}" ]]; then
  echo "==> judge"
  if [[ -z "$PROVIDER" ]]; then
    PROVIDER=$([[ -n "${ANTHROPIC_API_KEY:-}" ]] && echo "anthropic" || echo "noop")
  fi
  go -C tool run ./cmd/genspecs judge \
    --provider "$PROVIDER" \
    --specs-json products/specs.json \
    --harness-json products/harness_signals.json \
    --sdk-repos-root "$ROOT" \
    --specs-repo "$ROOT/sdk-specs" \
    --out products/spec_support.json \
    "${EXTRA_ARGS[@]}"
fi

echo "==> html"
go -C tool run ./cmd/genspecs html \
  --specs-json products/specs.json \
  --support-json products/spec_support.json \
  --out-dir _site

echo "Done. Artifacts:"
echo "  products/specs.json"
echo "  products/harness_signals.json"
echo "  products/spec_support.json"
echo "  _site/spec-support.html"
echo "  _site/spec-support-by-sdk.html"
