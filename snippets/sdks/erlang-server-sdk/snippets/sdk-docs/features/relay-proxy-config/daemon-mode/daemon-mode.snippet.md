---
id: erlang-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Daemon mode configuration example for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
ldclient:start_instance("YOUR_SDK_KEY", #{
  use_ldd => true,
  feature_store => your_feature_store
})
```
