---
id: erlang-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Service endpoint configuration example for Erlang.
---

```erlang
ldclient:start_instance("YOUR_SDK_KEY", #{
  stream_uri => "https://stream.launchdarkly.us",
  base_uri => "https://sdk.launchdarkly.us",
  events_uri => "https://events.launchdarkly.us"
})
```
