---
id: erlang-server-sdk/sdk-docs/features/anonymous/anonymous-v1x
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Anonymous user example for Erlang, SDK v1.x.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
Key = <<"example-user-key">>,
Anonymous = true,
Map = #{
    key => Key,
    anonymous => Anonymous
},
User = ldclient_user:new_from_map(Map)
```
