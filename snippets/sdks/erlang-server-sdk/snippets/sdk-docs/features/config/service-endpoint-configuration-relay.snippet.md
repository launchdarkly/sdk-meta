---
id: erlang-server-sdk/sdk-docs/features/config/service-endpoint-configuration-relay
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Service endpoint configuration example for Erlang.
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
