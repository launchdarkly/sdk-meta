---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-working-with-built-in-and-custom-attributes-1-x-syntax-user-with-attributes
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "1.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```erlang
User = ldclient_user:set(first_name, <<"Sandy">>,
  ldclient_user:set(last_name, <<"Smith">>,
  ldclient_user:set(email, <<"sandy@example.com">>,
  ldclient_user:new(<<"example-user-key">>))))

%% Or as a map

User = #{
  key => <<"example-user-key">>,
  first_name => <<"Sandy">>,
  last_name => <<"Smith">>,
  email => <<"sandy@example.com">>
},


```
