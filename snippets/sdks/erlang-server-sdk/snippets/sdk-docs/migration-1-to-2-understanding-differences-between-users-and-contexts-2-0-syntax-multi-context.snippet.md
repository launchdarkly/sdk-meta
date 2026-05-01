---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-2-0-syntax-multi-context
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "2.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```erlang
Context = ldclient_context:new_multi_from([
  %% Using `new/1` creates a context with a kind of <<"user">>.
  ldclient_context:new(<<"example-user-key">>),
  %% Using `new/2` creates a context of the specified kind (<<"device">>).
  ldclient_context:new(<<"example-device-key">>, <<"device">>)]). %% kind = device
```
