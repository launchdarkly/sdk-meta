---
id: erlang-server-sdk/sdk-docs/evaluate-a-context-erlang-sdk-v1-x
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Erlang SDK v1.x in section \"Evaluate a context\""
# Bucket C: erlang-server validator's gen_server harness is incompatible with the erlang-syntax-only scaffold's module shape. See _sdk-docs-port-notes.md.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
Flag = ldclient:variation(<<"example-flag-key">>, #{key => <<"example-user-key">>}, false)
```
