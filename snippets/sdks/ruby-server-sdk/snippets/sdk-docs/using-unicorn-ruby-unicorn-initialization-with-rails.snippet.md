---
id: ruby-server-sdk/sdk-docs/using-unicorn-ruby-unicorn-initialization-with-rails
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Unicorn initialization with Rails in section \"Using Unicorn\""
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
after_fork do |server,worker|
  Rails.configuration.client.postfork
end
```
