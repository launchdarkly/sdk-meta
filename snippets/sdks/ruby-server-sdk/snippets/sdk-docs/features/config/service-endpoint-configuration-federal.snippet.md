---
id: ruby-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Service endpoint configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
config = LaunchDarkly::Config.new(
  stream_uri: "https://stream.launchdarkly.us",
  base_uri: "https://sdk.launchdarkly.us",
  events_uri: "https://events.launchdarkly.us")
```
