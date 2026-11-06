---
id: ruby-server-sdk/sdk-docs/features/overrides/override-affected
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Check whether a flag override affected an evaluation in Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
detail = client.variation_detail("my-boolean-flag-key", context, false)
if detail.reason.override_affected
  puts "An override affected this evaluation"
end
```
