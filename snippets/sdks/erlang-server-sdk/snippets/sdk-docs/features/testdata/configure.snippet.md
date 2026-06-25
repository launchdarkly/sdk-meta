---
id: erlang-server-sdk/sdk-docs/features/testdata/configure
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Test data source configuration for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only-block
---

```erlang
Options = #{
  datasource => testdata,
  send_events => false,
  feature_store => ldclient_storage_map
},
ldclient:start_instance(SdkKey, Options),
```
