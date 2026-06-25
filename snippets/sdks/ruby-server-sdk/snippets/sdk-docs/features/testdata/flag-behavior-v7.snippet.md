---
id: ruby-server-sdk/sdk-docs/features/testdata/flag-behavior-v7
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Configuring test data flag behavior for Ruby SDK v7.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
# This flag is true for the context key "example-context-key" and false for everyone else
td.update(td.flag("flag-key-456def").
    variation_for_user("example-context-key", true).
    fallthrough_variation(false))

# This flag returns the string variation "green" for contexts who have the custom
# attribute "admin" with a value of true, and "red" for everyone else.
td.update(td.flag("flag-key-789ghi").
    variations("red", "green").
    fallthrough_variation(0).
    if_match_context("user", "admin", true).then_return(1))
```
