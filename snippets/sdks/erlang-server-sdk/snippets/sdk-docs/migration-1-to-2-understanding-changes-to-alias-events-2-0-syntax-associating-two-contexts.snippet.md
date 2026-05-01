---
id: erlang-server-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-alias-events-2-0-syntax-associating-two-contexts
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "2.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```erlang
Context = ldclient_context:new_multi_from([
  ldclient_context:new(<<"example-user-key">>),
  ldclient_context:new(<<"example-device-key">>, <<"device">>)]),
ldclient_identify(Context).
```
