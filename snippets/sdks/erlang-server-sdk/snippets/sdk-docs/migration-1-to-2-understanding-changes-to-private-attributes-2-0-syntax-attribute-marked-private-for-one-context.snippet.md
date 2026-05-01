---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-private-attributes-2-0-syntax-attribute-marked-private-for-one-context
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "2.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```erlang
ContextWithPrivateAttributes = ldclient_context:set_private_attributes([<<"name">>, <<"/address/street">>],
    ldclient_context:set(<<"name">>, <<"Global Health Services">>,
    ldclient_context:set(<<"email">>, <<"info@globalhealthexample.com">>,
    ldclient_context:set(<<"address">>, #{
        <<"street">> => <<"123 Main Street">>,
        <<"city">> => <<"Springfield">>
    },
ldclient_context:new(<<"context-key-456def">>, <<"organization">>))))),
```
