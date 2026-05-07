---
id: ruby-server-sdk/sdk-docs/using-puma-ruby-puma-initialization-with-rails
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Puma initialization with Rails in section \"Using Puma\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
on_worker_boot do
  Rails.configuration.client.postfork
end
```
