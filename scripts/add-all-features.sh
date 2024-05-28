#!/bin/bash

set -e

# Given a csv that has format (sdk,featureName,versionIntroduced), we want to call
# the script ./scripts/add-feature.sh with teh foliwing arguments:
# (1) The SDK name. This will be provided as a command line argument (1).
# (2) The feature name. This comes from teh second column of the CSV.
# (3) The version introduced. This comes from the third column of the CSV.

# Since the CSV contains many rows, we should first grep it by argument (1) and then
# pipe the results (maybe to xargs or something?) in order to invoke the ./scripts/add-feature.sh
# repeatedly.

# Assume hte csv location is hardcoded as ./all.csv

# Example invocation: ./scripts/add-feature cpp-client-sdk multiEnv 3.1.0

#sdk=$1
csv=./all.csv

# If first argument is -create-pr, then set a bool to true
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

#if [ -z "$sdk" ]; then
#  echo "Usage: $0 <sdk>"
#  exit 1
#fi

#
#sdk_name=$(echo "$sdk" | cut -d'/' -f2)
#
#echo "SDK Name: $sdk_name"
#echo "REPO: $sdk"
## Now to

# Get list of SDK IDs from ./products/names.json. The data is an object with keys, each key
# is an SDK ID. Use jq. Then for each ID, run the code below:

ids=$(jq -r 'keys[]' ./products/names.json)

for id in $ids; do
  repo_link=$(jq -r ".\"$id\".github" ./products/repos.json)
  # The linke consists of launchdarkly/sdk-id. Strip off the launchdarkly/ part to get
  # the repo only:
  repo=$(echo "$repo_link" | cut -d'/' -f2)
  grep "$id" "$csv" | while IFS=, read -r sdk_id feature version; do
    echo "Adding feature $feature to $repo at version $version"
    ./scripts/add-feature.sh "$repo/$id" "$feature" "$version"
  done
done

GH_USERNAME=$(gh api user | jq -r .login)

if [ "$create_prs" = true ]; then
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

if [ "$copy_metadata" = true ]; then
  mkdir -p ./backfill-features
  echo "Copying metadata"
  for id in $ids; do
    repo_link=$(jq -r ".\"$id\".github" ./products/repos.json)
    # The linke consists of launchdarkly/sdk-id. Strip off the launchdarkly/ part to get
    # the repo only:
    repo=$(echo "$repo_link" | cut -d'/' -f2)

    cp "$repo/.sdk_metadata.json" "./backfill-features/launchdarkly_$repo.json"
  done
fi
