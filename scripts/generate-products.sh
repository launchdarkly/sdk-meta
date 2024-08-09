#!/bin/bash

set -e

# Notes:
# -S argument to JQ is used to sort the keys of the output objects so we get more deterministic output,
# and it's easier to compare diffs between commits to the repo.

# Cleanup existing products so we have a clean slate.
rm products/*.json
rm api/sdkmeta/data/*.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_languages;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] += [$item.language])' > products/languages.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_names;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] = $item.name)' > products/names.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_types;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] = $item.type)' > products/types.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_repos;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] += {github: $item.github})' > products/repos.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_features;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] += {($item.feature): {introduced: $item.introduced, deprecated: $item.deprecated, removed: $item.removed}})' > products/features.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_feature_info;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] += {name: $item.name, description: $item.description})' > products/feature_info.json

./scripts/eols.sh metadata.sqlite3  |
  jq -n 'reduce inputs[] as $input ({}; .[$input.id] += [$input | del(.id)])' > products/releases.json

# JSON products are duplicated into a subdirectory of the API Go module, because the 'embed' package only allows
# us to embed files that are in the same directory tree as the module.
cp products/*.json api/sdkmeta/data/
