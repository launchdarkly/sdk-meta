#!/bin/bash

set -e


if [ -z "$1" ]; then
  echo "Usage: $0 <sqlite-db-path> <metadata-dir-path>"
  exit 1
fi

go build ./cmd/ingest

temp_db=$1
rm -f "$temp_db"

temp_dir=$2
rm -rf "$temp_dir"

sqlite3 "$temp_db" < ./schemas/sdk_metadata.sql
mkdir "$temp_dir"

jq -r '.repos[]' < config.json | while read -r repo; do
  echo "Fetching metadata.json for $repo"
  sanitized_repo=$(echo "$repo" | tr '/' '_')
  gh api "repos/$repo/contents/.sdk_metadata.json" -q '.content' | base64 --decode > "$temp_dir/$sanitized_repo.json"
  echo "Ingesting metadata.json for $repo"
  ./ingest -metadata "$temp_dir/$sanitized_repo.json" -db "$temp_db" -repo "$repo"
done
