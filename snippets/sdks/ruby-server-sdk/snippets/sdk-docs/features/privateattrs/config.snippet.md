---
id: ruby-server-sdk/sdk-docs/features/privateattrs/config
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Private attribute configuration for Ruby SDK v7.0.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
# All attributes marked private
config = LaunchDarkly::Config.new({all_attributes_private: true})

# Two attributes marked private
config = LaunchDarkly::Config.new({private_attributes: ["name", "email"]})
```
