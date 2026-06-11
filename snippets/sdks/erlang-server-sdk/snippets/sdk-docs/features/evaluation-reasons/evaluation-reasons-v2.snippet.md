---
id: erlang-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v2
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Flag evaluation reason example for Erlang SDK v2.0+.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
Flag = ldclient:variation_detail(<<"example-flag-key">>, #{key => <<"example-context-key">>}, false)
```
