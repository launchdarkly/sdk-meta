#!/bin/bash

for file in ./products/*.json; do
  jq . "$file" > "$file.tmp" && mv "$file.tmp" "$file"
done
