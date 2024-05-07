#!/bin/bash

set -e

# This script can be used to add a feature (or deprecate one) for a particular SDK.
# For example, using a monorepo: ./add-feature.sh cpp-sdks/cpp-client-sdk appMetadata 3.0
# Or a solorepo: ./add-feature.sh ruby-server-sdk appMetadata 3.0
repo=$1
feature=$2
introduced=$3

if [ -z "$repo" ] || [ -z "$feature" ] || [ -z "$introduced" ]; then
  echo "Usage: $0 <repo> <feature> <introduced>"
  exit 1
fi


sqlite3 metadata.sqlite3 "SELECT COUNT(*) FROM sdk_feature_info WHERE id = '$feature';" |
  grep 1 || {
    echo "Feature '$feature' not recognized. Check spelling, or add it to the metadata.sqlite3 database first.";
    echo "Known features:"
    sqlite3 metadata.sqlite3 "SELECT id FROM sdk_feature_info;"
    exit 1;
  }

# The repo name is either going to be something like ruby-server-sdk,
# or a SDK within a monorepo, like js-core/node-server-sdk.
# If there's a slash, we need to split it into the two components so that we can clone the first one
# in both cases.

if [[ "$repo" == *"/"* ]]; then
  repo_name=$(echo "$repo" | cut -d'/' -f1)
else
  repo_name=$repo
fi

gh repo clone "launchdarkly/$repo" -- --depth=1

if [ ! -f "$repo_name/.sdk_metadata.json" ]; then
  echo ".sdk_metadata.json not found, run ./scripts/add-repo.sh first"
  exit 1
fi

# Now, if the
