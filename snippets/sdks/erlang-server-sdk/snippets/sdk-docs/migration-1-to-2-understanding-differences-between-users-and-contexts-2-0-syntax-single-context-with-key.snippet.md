---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-2-0-syntax-single-context-with-key
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "2.0 syntax, single context with key in section \"Understanding differences between users and contexts\""
---

```erlang
Context = ldclient_context:new(<<"example-organization-key">>, <<"organization">>)
%% Or as a map
Context = #{kind => <<"organization">>, key => <<"example-organization-key">>}
```
