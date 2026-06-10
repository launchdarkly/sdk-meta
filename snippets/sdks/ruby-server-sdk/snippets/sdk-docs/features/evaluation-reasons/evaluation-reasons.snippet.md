---
id: ruby-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Flag evaluation reason example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
detail = client.variation_detail("example-flag-key", my_context, false)
value = detail.value
index = detail.variation_index
reason = detail.reason
```
