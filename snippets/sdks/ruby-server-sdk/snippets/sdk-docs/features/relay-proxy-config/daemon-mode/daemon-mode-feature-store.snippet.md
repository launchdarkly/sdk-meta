---
id: ruby-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-feature-store
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Daemon mode configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby

store = SomeKindOfFeatureStore.new(storeOptions)

config = LaunchDarkly::Config.new(
  feature_store: store,
  use_ldd: true
)
```
