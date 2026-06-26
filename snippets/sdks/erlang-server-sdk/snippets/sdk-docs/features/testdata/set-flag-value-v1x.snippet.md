---
id: erlang-server-sdk/sdk-docs/features/testdata/set-flag-value-v1x
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Setting a test data flag to a specific value for Erlang SDK v1.x.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only-block
---

```erlang
{ok, Flag} = ldclient_testdata:flag("example-flag-key"),
ldclient_testdata:update(ldclient_flagbuilder:variation_for_all_users(true, Flag)),
```
