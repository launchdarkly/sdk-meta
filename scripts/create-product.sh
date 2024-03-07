#!/bin/bash

# Usage: ./new-data-template.sh table-name
# Outputs a new JSON file in /data with the given name.

if [ -z "$1" ]; then
  echo "Usage: $0 <table-name>"
  exit 1
fi

jq 'map({(.): {}}) | add' data/sdks.json > data/"$1".json
