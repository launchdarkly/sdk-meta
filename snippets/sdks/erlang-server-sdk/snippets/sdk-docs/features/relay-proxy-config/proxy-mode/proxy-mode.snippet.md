---
id: erlang-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Proxy mode configuration example for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
ldclient:start_instance("YOUR_SDK_KEY", #{
  stream_uri => "https://your-relay-proxy.com:8030",
  base_uri => "https://your-relay-proxy.com:8030",
  events_uri => "https://your-relay-proxy.com:8030"
})
```
