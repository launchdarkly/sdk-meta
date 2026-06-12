---
id: ruby-server-sdk/sdk-docs/features/testdata/configure
sdk: ruby-server-sdk
kind: reference
lang: ruby
description: Test data source configuration for Ruby.
validation:
  scaffold: ruby-server-sdk/scaffolds/ruby-syntax-only
---

```ruby
require 'ldclient-rb'

td = LaunchDarkly::Integrations::TestData.data_source
# You can set any initial flag states here with td.update

config = LaunchDarkly::Config.new(data_source: td)
client = LaunchDarkly::LDClient.new(sdk_key, config)
```
