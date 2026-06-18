---
id: erlang-server-sdk/sdk-docs/features/contextconfig/context-example-v2
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Context example for Erlang SDK v2.0+.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
    Context = ldclient_context:set(<<"name">>, <<"Sandy Smith">>,
        ldclient_context:set(<<"email">>, <<"sandy@example.com">>,
        ldclient_context:set(<<"group">>, [<<"microsoft">>, <<"google">>],
        ldclient_context:new(<<"example-user-key">>, <<"user">>))))
```
