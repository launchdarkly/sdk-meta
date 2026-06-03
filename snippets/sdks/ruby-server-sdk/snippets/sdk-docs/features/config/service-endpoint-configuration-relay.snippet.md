---
id: ruby-server-sdk/sdk-docs/features/config/service-endpoint-configuration-relay
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Service endpoint configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
config = LaunchDarkly::Config.new(
  stream_uri: "https://your-relay-proxy.com:8030",
  base_uri: "https://your-relay-proxy.com:8030",
  events_uri: "https://your-relay-proxy.com:8030")
```
