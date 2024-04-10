#!/bin/bash

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

ID="$ID" NAME="$NAME" TYPE="$TYPE" LANG="$LANG" envsubst < ""./scripts/metadata-template.json > "$repo_name/.sdk_metadata.json"
(
  cd "$repo_name" || exit
  git switch -c cw/add-sdk-metadata
  git add .sdk_metadata.json
  git commit -m "chore: add .sdk_metadata.json"
  gh pr create --fill
)
