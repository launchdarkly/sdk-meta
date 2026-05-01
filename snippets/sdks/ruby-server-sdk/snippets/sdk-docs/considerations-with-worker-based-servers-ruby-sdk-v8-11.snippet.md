---
id: ruby-server-sdk/sdk-docs/considerations-with-worker-based-servers-ruby-sdk-v8-11
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby SDK v8.11+ in section \"Considerations with worker-based servers\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
# 1. Create the client before forking.
client = LaunchDarkly::LDClient.new("YOUR_SDK_KEY")

# 2. From the newly forked process, reinitialize the client by calling `postfork`.
# Examples for specific servers are shown below.
client.postfork
```
