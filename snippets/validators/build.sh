#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
    echo "usage: $0 <runtime>" >&2
    exit 2
fi

validators_dir=$(cd -- "$(dirname -- "$0")" && pwd)
runtime=$1
dockerfile="$validators_dir/languages/$runtime/Dockerfile"
runner="$validators_dir/languages/$runtime/runner.yaml"

if [[ ! -f "$dockerfile" ]]; then
    echo "validator Dockerfile not found: $dockerfile" >&2
    exit 1
fi
if [[ ! -f "$runner" ]]; then
    echo "validator runner configuration not found: $runner" >&2
    exit 1
fi

image_prefix=$(sed -nE 's/^image-prefix:[[:space:]]*//p' "$runner")
if [[ -z "$image_prefix" ]]; then
    echo "validator runner has no image-prefix: $runner" >&2
    exit 1
fi

# The snippets CLI resolves floating SDK families in shared/versions/sdks.env
# to exact releases, matching what `snippets validate` builds with.
build_args=()
while IFS= read -r arg; do
    build_args+=(--build-arg "$arg")
done < <(cd "$validators_dir/.." && go run ./cmd/snippets validator-build-args --validators="$validators_dir" "$runtime")

docker build --progress=plain \
    -f "$dockerfile" \
    "${build_args[@]}" \
    -t "$image_prefix" \
    "$validators_dir"
