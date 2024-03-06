#!/bin/bash

set -e

mapfile -t schemas < <(git grep -Fl 'https://json-schema.org' '**/*.json')

function runTest() {
    primary=$(realpath --relative-to="$(pwd)" "$1")

    # Build up a list of reference schemas to use, being careful to ignore the
    # schema under test. If we don't skip that one, ajv will generate an error
    # about duplicate $ids.
    ajvFlags=()
    for schema in "${schemas[@]}"; do
        if [[ "$schema" != "$primary" ]]; then
            ajvFlags+=(-r "${schema}")
        fi
    done

    npx --package=ajv-cli --package=ajv-formats ajv validate --spec=draft2020 -s "$primary" "${ajvFlags[@]}" -d "$2"
}

runTest ./schemas/features.json ./data/features.json
runTest ./schemas/types.json ./data/types.json
