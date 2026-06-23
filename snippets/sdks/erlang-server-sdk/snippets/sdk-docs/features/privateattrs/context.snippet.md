---
id: erlang-server-sdk/sdk-docs/features/privateattrs/context
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Marking context attributes private for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
    ContextWithPrivateAttributes = ldclient_context:set_private_attributes([<<"name">>, <<"/address/street">>],
        ldclient_context:set(<<"name">>, <<"Global Health Services">>,
        ldclient_context:set(<<"email">>, <<"info@globalhealthexample.com">>,
        ldclient_context:set(<<"address">>, #{
            <<"street">> => <<"123 Main Street">>,
            <<"city">> => <<"Springfield">>
        },
    ldclient_context:new(<<"context-key-456def">>, <<"organization">>)))))
```
