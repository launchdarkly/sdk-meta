---
id: erlang-server-sdk/sdk-docs/features/config/index
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: SDK configuration example for Erlang.

---

```erlang
% Specify options
ldclient:start_instance("YOUR_SDK_KEY", #{stream => false})

% With a custom instance name
ldclient:start_instance("YOUR_SDK_KEY", your_instance, #{stream => false})
```
