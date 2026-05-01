---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-2-0-syntax-context-with-key
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "2.0 syntax, context with key in section \"Understanding differences between users and contexts\""
---

```erlang
Context = ldclient_context:new(<<"example-user-key">>),
%% Or as a map
Context = #{key => <<"example-user-key">>, kind => <<"user">>}
```
