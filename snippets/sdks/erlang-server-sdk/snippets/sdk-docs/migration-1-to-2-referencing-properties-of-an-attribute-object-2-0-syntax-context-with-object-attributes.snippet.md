---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-referencing-properties-of-an-attribute-object-2-0-syntax-context-with-object-attributes
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "2.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```erlang
Context =
  ldclient_context:set(<<"address">>, #{
    <<"street">> => <<"123 Main Street">>,
    <<"city">> => <<"Springfield">>
  },
  ldclient_context:new(<<"context-key-456def">>, <<"organization">>)),
```
