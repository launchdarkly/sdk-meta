---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-1-x-syntax-user-with-key
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "1.x syntax, user with key in section \"Understanding differences between users and contexts\""
---

```erlang
User = ldclient_user:new(<<"example-user-key">>),
%% Or as a map
User = #{key => <<"example-user-key">>}
```
