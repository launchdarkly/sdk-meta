---
id: ruby-server-sdk/sdk-docs/features/filedata/flags-from-files
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: File data source configuration example for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only

---

```ruby
require 'ldclient-rb'

data_source = LaunchDarkly::Integrations::FileData.data_source(
  paths: [ "file1.json", "file2.json" ],
  auto_update: true
)

config = LaunchDarkly::Config.new(
  data_source: data_source,
  send_events: false
)

client = LaunchDarkly::LDClient.new("sdk key", config)
```
