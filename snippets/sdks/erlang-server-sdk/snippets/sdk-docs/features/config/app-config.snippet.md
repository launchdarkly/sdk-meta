---
id: erlang-server-sdk/sdk-docs/features/config/app-config
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Application metadata configuration example for Erlang.
---

```erlang

ldclient:start_instance("YOUR_SDK_KEY", #{
  application => #{
    id => <<"authentication-service">>,
    version => <<"1.0.0">>
  }
})

```
