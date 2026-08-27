---
id: erlang-server-sdk/sdk-info/init
sdk: erlang-server-sdk
kind: init
lang: erlang
file: erlang-server-sdk/init.txt
description: |
  Client initialization snippet for erlang-server-sdk. Compile-checked
  via the erlang-syntax-only scaffold (the runtime path requires the
  snippet to be the pre-baked project's hello_erlang_server gen_server,
  which is not a useful user-facing init snippet).
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
% This starts the SDK instance with the default options.
% Never hardcode your SDK key in production.
ldclient:start_instance("YOUR_SDK_KEY")
```
