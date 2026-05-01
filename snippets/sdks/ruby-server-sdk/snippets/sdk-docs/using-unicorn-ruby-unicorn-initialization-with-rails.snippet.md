---
id: ruby-server-sdk/sdk-docs/using-unicorn-ruby-unicorn-initialization-with-rails
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: "Ruby Unicorn initialization with Rails in section \"Using Unicorn\""
---

```ruby
after_fork do |server,worker|
  Rails.configuration.client.postfork
end
```
