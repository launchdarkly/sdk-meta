#!/bin/bash

# Checks that for each SDK ID defined in sdks.json, all other files have an entry for that ID.
# If not, outputs a diff containing the missing IDs and exits with a non-zero status.

cleanup() {
  rm ./data/*.json.keys
}

trap cleanup EXIT

jq -r 'values[]' ./data/sdks.json | sort > ./data/sdks.json.keys

for file in ./data/*.json; do
  # Skip sdks.json itself
  if [ "$file" == "./data/sdks.json" ]; then
    continue
  fi
  jq -r 'keys[]' "$file" | sort > "$file.keys"
  if diff -q ./data/sdks.json.keys "$file.keys" > /dev/null; then
    echo "$file is consistent"
  else
    echo "$file is missing some SDK IDs:"
    diff ./data/sdks.json.keys "$file.keys"
    exit 1
  fi
done
