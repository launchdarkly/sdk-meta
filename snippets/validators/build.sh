#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
    echo "usage: $0 <runtime>" >&2
    exit 2
fi

validators_dir=$(cd -- "$(dirname -- "$0")" && pwd)
runtime=$1
dockerfile="$validators_dir/languages/$runtime/Dockerfile"

if [[ ! -f "$dockerfile" ]]; then
    echo "validator Dockerfile not found: $dockerfile" >&2
    exit 1
fi

source "$validators_dir/shared/versions/images.env"
source "$validators_dir/shared/versions/npm.env"
source "$validators_dir/shared/versions/toolchains.env"

build_args=()
while IFS= read -r name; do
    [[ -z "$name" ]] && continue
    if [[ ! -v "$name" ]]; then
        echo "validator Dockerfile declares ARG $name but no version entry exists" >&2
        exit 1
    fi
    build_args+=(--build-arg "$name=${!name}")
done < <(sed -nE 's/^[[:space:]]*ARG[[:space:]]+([A-Za-z_][A-Za-z0-9_]*)(=.*)?[[:space:]]*$/\1/p' "$dockerfile" | sort -u)

docker build --progress=plain \
    -f "$dockerfile" \
    "${build_args[@]}" \
    -t "sdk-snippets/$runtime-validator" \
    "$validators_dir"
