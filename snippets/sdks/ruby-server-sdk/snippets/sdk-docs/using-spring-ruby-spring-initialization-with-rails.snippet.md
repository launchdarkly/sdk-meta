---
id: ruby-server-sdk/sdk-docs/using-spring-ruby-spring-initialization-with-rails
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Spring initialization with Rails in section \"Using Spring\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
Spring.after_fork do
  Rails.configuration.client.postfork
end
```
