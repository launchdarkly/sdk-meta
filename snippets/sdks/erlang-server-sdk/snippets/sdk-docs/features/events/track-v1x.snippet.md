---
id: erlang-server-sdk/sdk-docs/features/events/track-v1x
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Custom event tracking example for Erlang SDK v1.x.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
ldclient:track(<<"example-event-key">>, #{key => <<"example-user-key">>}, #{data => <<"example">>})
```
