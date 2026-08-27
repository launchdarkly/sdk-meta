---
id: ruby-server-sdk/sdk-info/init
sdk: ruby-server-sdk
kind: init
lang: ruby
file: ruby-server-sdk/init.txt
description: |
  Client initialization snippet for ruby-server-sdk. A complete
  runnable program modeled on launchdarkly/hello-ruby: construct the
  client (the constructor blocks until the connection is up or times
  out) and report success or failure. The wrappee owns the whole
  program, so the harness asserts the snippet's own success line via
  SNIPPET_SUCCESS_RE instead of wrapping it in a trailer-emitting
  scaffold.
validation:
  runtime: ruby
  requirements: |
    launchdarkly-server-sdk
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```ruby
require 'ldclient-rb'

# This is your LaunchDarkly SDK key.
# Never hardcode your SDK key in production.
client = LaunchDarkly::LDClient.new('YOUR_SDK_KEY')

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
