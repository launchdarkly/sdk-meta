---
id: ruby-server-sdk/sdk-docs/features/testdata/set-flag-value-v7
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Setting a test data flag to a specific value for Ruby SDK v7.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
td.update(td.flag("example-flag-key").variation_for_all(false))
```
