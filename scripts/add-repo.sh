#!/bin/bash

set -e

# This script can be used to create a basic metadata file for a particular repo.
# It will clone the repo, ask some questions, then commit the changes and
# push to Github.
#
# It's not very smart - ideally we'd have a JSON schema for the metadata
# file so we could validate it before committing. For now, it relies on humans (you) to make sure
# it is correct.

repo=$1

if [ -z "$repo" ]; then
  echo "Usage: $0 <repo>"
  exit 1
fi

gh repo clone "$repo" -- --depth=1
repo_name=$(basename "$repo")

if [ -f "$repo_name/.sdk_metadata.json" ]; then
  echo "metadata already exists for: $repo_name"
  exit 0
fi

echo "create metadata for: $repo_name"
echo "SDK ID: "
read -r ID
echo "SDK name: "
read -r NAME
echo "SDK type: "
read -r TYPE
echo "SDK language: "
read -r LANG

GH_USERNAME=$(gh api user | jq -r .login)

ID="$ID" NAME="$NAME" TYPE="$TYPE" LANG="$LANG" envsubst < ""./scripts/metadata-template.json > "$repo_name/.sdk_metadata.json"
(
  cd "$repo_name" || exit
  git switch -c "$GH_USERNAME"/add-sdk-metadata
  git add .sdk_metadata.json
  git commit -m "chore: add .sdk_metadata.json"
  gh pr create --fill
)
