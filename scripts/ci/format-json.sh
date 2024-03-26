#!/bin/bash

jq . config.json > config.json.tmp && mv config.json.tmp config.json

for file in ./products/*.json; do
  jq . "$file" > "$file.tmp" && mv "$file.tmp" "$file"
done
