---
id: erlang-server-sdk/sdk-docs/features/anonymous/anonymous-v2
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Anonymous context example for Erlang, SDK v2.0+.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
    Context = ldclient_context:set(anonymous, true,
        ldclient_context:new(<<"example-user-key">>))
```
