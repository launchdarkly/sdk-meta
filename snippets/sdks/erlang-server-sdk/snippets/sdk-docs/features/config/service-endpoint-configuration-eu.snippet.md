---
id: erlang-server-sdk/sdk-docs/features/config/service-endpoint-configuration-eu
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Service endpoint configuration example for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
ldclient:start_instance("YOUR_SDK_KEY", #{
  stream_uri => "https://stream.eu.launchdarkly.com",
  base_uri => "https://sdk.eu.launchdarkly.com",
  events_uri => "https://events.eu.launchdarkly.com"
})
```
