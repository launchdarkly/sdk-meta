---
id: ruby-server-sdk/sdk-docs/features/config/service-endpoint-configuration-eu
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Service endpoint configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
config = LaunchDarkly::Config.new(
  stream_uri: "https://stream.eu.launchdarkly.com",
  base_uri: "https://sdk.eu.launchdarkly.com",
  events_uri: "https://events.eu.launchdarkly.com")
```
