---
id: erlang-server-sdk/sdk-docs/features/evaluating/evaluating-v1x
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Flag evaluation example for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
Flag = ldclient:variation(<<"example-flag-key">>, #{key => <<"example-user-key">>}, false, your_instance)
```
