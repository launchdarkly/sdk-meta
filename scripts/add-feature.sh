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
  is_monorepo=true
else
  repo_name=$repo
  is_monorepo=false
fi

if [ -d "$repo_name" ]; then
  echo "note: repo already cloned, will update"
else
  gh repo clone "launchdarkly/$repo_name" -- --depth=1
fi

if [ ! -f "$repo_name/.sdk_metadata.json" ]; then
  echo ".sdk_metadata.json not found, run ./scripts/add-repo.sh first"
  exit 1
fi

# Now, if is_monorepo is true, then we need to access the only child SDK of the "sdks" key in the .sdk_metadata.json
# file. That is, the file will have an "sdks" : { "some-sdk-id" : { .. } } structure, and we need to automatically
# be editing the object some-sdk-id.
# Otherwise, if this is a monorepo, we need to edit the child object of "sdks" that matches the sdk name (the second
# part of the repo name after the slash.)
# What we'll do is add a new sub key, if it doesn't exist, named "features". Then we'll add the feature name
# that was specified in the argument. The value is an object with a single key, "introduced".
# Then we'll set "introduced"'s value to the version specified as an argument. If the feature
# key already exists, we'll update it's value to the new version and echo a warning.
# Here's how to do it in jq:

sdk_id=$(basename "$repo")

if [ "$is_monorepo" = true ]; then
  jq --arg feature "$feature" --arg introduced "$introduced" --arg sdk "$sdk_id" '.sdks[$sdk].features |= . + { ($feature): { introduced: $introduced } }' "$repo_name/.sdk_metadata.json" > "$repo_name/.sdk_metadata.json.tmp"
else
  jq --arg feature "$feature" --arg introduced "$introduced" '.sdks[.sdks | keys[0]].features |= . + { ($feature): { introduced: $introduced } }' "$repo_name/.sdk_metadata.json" > "$repo_name/.sdk_metadata.json.tmp"
fi

mv "$repo_name/.sdk_metadata.json.tmp" "$repo_name/.sdk_metadata.json"
