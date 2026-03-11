#!/bin/bash

set -e


if [ -z "$1" ]; then
  echo "Usage: $0 <sqlite-db-path> <metadata-dir-path>"
  exit 1
fi

(
  cd tool
  go build ./cmd/ingest
)

temp_db=$1
rm -f "$temp_db"

temp_dir=$2
rm -rf "$temp_dir"

sqlite3 "$temp_db" < ./schemas/sdk_metadata.sql
mkdir "$temp_dir"

# Phase 1: Old-repo backfills. These capture releases from repositories that are
# no longer the canonical home for an SDK (e.g. standalone repos that have been
# merged into a monorepo). They run first so their sdk_repos entries have lower
# rowids, which matters because generate-products.sh uses jq reduce (last writer
# wins) to build repos.json.
for file in ./backfill/*.json; do
  same_repo=$(jq -r '."same-repo" // false' "$file")
  if [ "$same_repo" = "true" ]; then
    continue
  fi
  repo=$(basename "$file" .json | tr '_' '/')
  echo "backfilling $repo"
  ./tool/ingest -metadata "$file" -db "$temp_db" -repo "$repo"
done

# Phase 2: Same-repo backfills. These capture old unprefixed releases from repos
# that have transitioned to prefixed tags but are still the canonical home for the
# SDK. They run after old-repo backfills so their sdk_repos entries take precedence
# in repos.json.
for file in ./backfill/*.json; do
  same_repo=$(jq -r '."same-repo" // false' "$file")
  if [ "$same_repo" != "true" ]; then
    continue
  fi
  repo=$(basename "$file" .json | tr '_' '/')
  echo "backfilling $repo"
  ./tool/ingest -metadata "$file" -db "$temp_db" -repo "$repo"
done

# Phase 3: Main crawl. Fetch live metadata and prefixed releases from GitHub.
# Repos with old-repo backfills are skipped (their metadata lives elsewhere now).
# Repos with same-repo backfills are still crawled for metadata and new releases.
./scripts/repos.sh | while read -r repo; do
  echo "checking $repo"
  sanitized_repo=$(echo "$repo" | tr '/' '_')

  if [ -f "./backfill/$sanitized_repo.json" ]; then
    same_repo=$(jq -r '."same-repo" // false' "./backfill/$sanitized_repo.json")
    if [ "$same_repo" != "true" ]; then
      echo "skipping $repo, it was backfilled"
      continue
    fi
  fi

  metadata=$(gh api "repos/$repo/contents/.sdk_metadata.json" -q '.content') || {
    continue
  }
  echo "$metadata" | base64 --decode > "$temp_dir/$sanitized_repo.json"
  echo "found metadata in $repo"
  ./tool/ingest -metadata "$temp_dir/$sanitized_repo.json" -db "$temp_db" -repo "$repo"
done
