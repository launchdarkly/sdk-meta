---
id: erlang-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Persistent feature store configuration example for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only-stmts

---

```erlang
LdOptions = #{
  feature_store => your_feature_store
},
ldclient:start_instance("YOUR_SDK_KEY", LdOptions).
```
