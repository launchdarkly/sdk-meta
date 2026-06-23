---
id: erlang-server-sdk/sdk-docs/features/contextconfig/multi-context-v2
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Multi-context example for Erlang SDK v2.0+.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
Context = ldclient_context:new_multi_from([
  %% Using `new/1` creates a context with a kind of <<"user">>.
  ldclient_context:new(<<"example-user-key">>),
  %% Using `new/2` creates a context of the specified kind (<<"device">>).
  ldclient_context:new(<<"example-device-key">>, <<"device">>)])
```
