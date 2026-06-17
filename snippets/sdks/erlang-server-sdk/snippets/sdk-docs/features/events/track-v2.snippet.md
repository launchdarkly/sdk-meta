---
id: erlang-server-sdk/sdk-docs/features/events/track-v2
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Custom event tracking example for Erlang SDK v2.0+.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
ldclient:track(<<"example-event-key">>, #{key => <<"example-context-key">>}, #{data => <<"example">>})
```
