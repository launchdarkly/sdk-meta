---
id: ruby-server-sdk/sdk-docs/evaluate-a-context-ruby-sdk-v7-0
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby SDK v7.0+ in section \"Evaluate a context\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
context = LaunchDarkly::LDContext.with_key("example-context-key")
show_feature = client.variation("example-flag-key", context, false)
if show_feature
  # application code to show the feature
else
  # the code to run if the feature is off
end
```
