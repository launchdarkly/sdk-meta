---
id: erlang-server-sdk/sdk-docs/features/contextconfig/context-kind-v2
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Context with a non-user kind for Erlang SDK v2.0+.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
Context = ldclient_context:new(<<"example-organization-key">>, <<"organization">>),
%% Or as a map
Context = #{kind => <<"organization">>, key => <<"example-organization-key">>}
```
