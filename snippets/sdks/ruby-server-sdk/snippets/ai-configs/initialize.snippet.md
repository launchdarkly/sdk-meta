---
id: ruby-server-sdk/ai-configs/initialize
sdk: ruby-server-sdk
kind: initialize
lang: ruby
file: ruby-server-sdk/ai-configs/initialize.txt
description: Initialize the LaunchDarkly client and AI Configs client for ruby-server-sdk.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
ld_client = LaunchDarkly::LDClient.new("{{sdkkey}}")
ai_client = LaunchDarkly::AI::LDAIClient.new(ld_client)
```
