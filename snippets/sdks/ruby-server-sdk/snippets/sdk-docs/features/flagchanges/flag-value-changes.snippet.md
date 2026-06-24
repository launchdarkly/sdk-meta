---
id: ruby-server-sdk/sdk-docs/features/flagchanges/flag-value-changes
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Flag value change subscription example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only-block

---

```ruby
class Listener
    def update(changed)
        puts "Flag #{changed.key} has changed from #{changed.old_value} to #{changed.new_value}"
    end
end

client.flag_tracker.add_flag_value_change_listener("example-flag-key", context, Listener.new)
```
