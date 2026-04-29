---
id: ruby-server-sdk/getting-started/install
sdk: ruby-server-sdk
kind: install
lang: shell
description: Add the SDK gem to the Gemfile and install with bundler.
inputs:
  version:
    type: string
    description: Optional pinned SDK version; when empty the pin is omitted.
    runtime-default: ""
ld-application:
  slot: install
---

Next, install the SDK:

```shell
echo "gem 'launchdarkly-server-sdk'{{ if version }}, '{{ version }}'{{ end }}" >> Gemfile && bundle install
```
