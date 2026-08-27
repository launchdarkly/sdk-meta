---
id: ruby-server-sdk/sdk-info/init-env
sdk: ruby-server-sdk
kind: init
lang: ruby
file: ruby-server-sdk/init-env.txt
description: |
  Client initialization snippet for ruby-server-sdk reading the SDK
  key from an environment variable. No `placeholders:` needed — the
  harness already exports LAUNCHDARKLY_SDK_KEY into the validate run.
validation:
  runtime: ruby
  requirements: |
    launchdarkly-server-sdk
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```ruby
require 'ldclient-rb'

# Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
# environment variable, so one build can run in every environment.
client = LaunchDarkly::LDClient.new(ENV['LAUNCHDARKLY_SDK_KEY'])

if client.initialized?
  puts 'SDK successfully initialized'
else
  puts 'SDK failed to initialize'
  exit 1
end

# For onboarding purposes only we flush events as soon as
# possible so we quickly detect your connection.
# You don't have to do this in practice because events are automatically flushed.
client.flush
```
