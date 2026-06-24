---
id: ruby-server-sdk/sdk-docs/features/flagchanges/flag-changes
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Flag change subscription example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only-block

---

```ruby
class Listener
    def update(status)
        puts "Flag #{status.key} has changed"
    end
end

client.flag_tracker.add_listener(Listener.new)
```
