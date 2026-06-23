---
id: erlang-server-sdk/sdk-docs/features/contextconfig/context-example-v1
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: User example for Erlang SDK v1.x.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
Map = #{
    key => <<"example-user-key">>,
    secondary => <<"secondary-123abc">>,
    ip => <<"198.51.100.0">>,
    country => <<"Canada">>,
    email => <<"sandy@example.com">>,
    first_name => <<"Sandy">>,
    last_name => <<"Smith">>,
    avatar => <<"avatar-123abc">>,
    name => <<"Sandy Smith">>,
    anonymous => false,
    <<"custom-attr-name">> => <<"custom-value">>
},
User = ldclient_user:new_from_map(Map)
```
