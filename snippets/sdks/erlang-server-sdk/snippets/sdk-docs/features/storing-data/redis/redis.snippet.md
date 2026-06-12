---
id: erlang-server-sdk/sdk-docs/features/storing-data/redis/redis
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Redis feature store configuration example for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only-stmts

---

```erlang
LdOptions = #{
  redis_host     => "redis",
  redis_port     => "6379",
  redis_prefix   => "default",
  feature_store  => ldclient_storage_redis,
  cache_ttl      => 15
},
ldclient:start_instance("YOUR_SDK_KEY", LdOptions).
```
