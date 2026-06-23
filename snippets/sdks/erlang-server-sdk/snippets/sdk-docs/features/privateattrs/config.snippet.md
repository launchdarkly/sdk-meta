---
id: erlang-server-sdk/sdk-docs/features/privateattrs/config
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: Private attribute configuration for Erlang.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only

---

```erlang
%% All attributes marked as private
ldclient:start_instance("YOUR_SDK_KEY", #{private_attributes => all}),

%% Two attributes marked as private
ldclient:start_instance("YOUR_SDK_KEY", #{private_attributes => [<<"email">>, <<"address">>]})
```
