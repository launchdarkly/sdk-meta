---
id: erlang-server-sdk/sdk-docs/features/shutdown/shutdown
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Client shutdown example for Erlang.

---

```erlang
ldclient:stop_all_instances()

% Stops the default instance
ldclient:stop_instance()

% Stops a named instance
ldclient:stop_instance(my_instance)
```
