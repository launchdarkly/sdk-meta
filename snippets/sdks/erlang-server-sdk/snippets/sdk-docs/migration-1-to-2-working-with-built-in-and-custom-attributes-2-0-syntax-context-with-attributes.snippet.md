---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-working-with-built-in-and-custom-attributes-2-0-syntax-context-with-attributes
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "2.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```erlang
%% Contexts do not do any casing/binary conversion for atoms like users did.
%% So first_name will not match the firstName attribute.
User = ldclient_context:set(<<"firstName">>, <<"Sandy">>,
  ldclient_context:set(<<"lastName">>, <<"Smith">>,
  ldclient_context:set(<<"email">>, <<"sandy@example.com">>,
  ldclient_context:new(<<"example-user-key">>))))

%% Or as a map

User = #{
  key => <<"example-user-key">>,
  kind => <<"user">>,
  attributes => #{
    <<"firstName">> => <<"Sandy">>,
    <<"lastName">> => <<"Smith">>,
    <<"email">> => <<"sandy@example.com">>
  }
}
```
