#!/bin/bash

set -e

csv=./all.csv

create_prs=false
copy_metadata=false

if [ "$1" == "-create-prs" ]; then
  create_prs=true
  shift
fi

if [ "$1" == "-copy-metadata" ]; then
  copy_metadata=true
  shift
fi


ids=$(jq -r 'keys[]' ./products/names.json | grep -v relay-proxy)

for id in $ids; do
  repo_link=$(jq -r ".\"$id\".github" ./products/repos.json)
  repo=$(echo "$repo_link" | cut -d'/' -f2)
  grep "$id" "$csv" | while IFS=, read -r sdk_id feature version; do
    echo "Adding feature $feature to $repo at version $version"
    ./scripts/add-feature.sh "$repo/$id" "$feature" "$version"
  done
done

GH_USERNAME=$(gh api user | jq -r .login)

if [ "$create_prs" == true ]; then
  echo "Creating PRs"
  for id in $ids; do
    repo_link=$(jq -r ".\"$id\".github" ./products/repos.json)
    # The linke consists of launchdarkly/sdk-id. Strip off the launchdarkly/ part to get
    # the repo only:
    repo=$(echo "$repo_link" | cut -d'/' -f2)

    (
      cd "$repo" || exit
      git switch -c "$GH_USERNAME"/update-version-metadata
      git add .sdk_metadata.json
      git commit -m "chore: update .sdk_metadata.json with feature table"
      gh pr create --fill
    )


  done
fi

if [ "$copy_metadata" == true ]; then
  mkdir -p ./backfill-features
  echo "Copying metadata"
  for id in $ids; do
    repo_link=$(jq -r ".\"$id\".github" ./products/repos.json)

    repo=$(echo "$repo_link" | cut -d'/' -f2)

    cp "$repo/.sdk_metadata.json" "./backfill-features/launchdarkly_$repo.json"
  done
fi
