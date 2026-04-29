---
id: erlang-server-sdk/getting-started/run-call
sdk: erlang-server-sdk
kind: run
lang: erlang
description: Erlang shell call to evaluate the flag.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered command.
ld-application:
  slot: run-call
---

Inside the rebar3 shell, call:

```erlang
hello_erlang_server:get(<<"{{ featureKey }}">>, "FALLBACK_VALUE", <<"user@example.com">>).
```
