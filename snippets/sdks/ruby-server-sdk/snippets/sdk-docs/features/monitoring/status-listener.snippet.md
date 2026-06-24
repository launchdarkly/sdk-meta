---
id: ruby-server-sdk/sdk-docs/features/monitoring/status-listener
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Data source status change listener for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
# listener#update will be called when the status is changed
client.data_source_status_provider.add_listener(listener);
```
