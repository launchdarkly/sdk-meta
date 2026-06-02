---
id: erlang-server-sdk/sdk-docs/initialize-the-client-erlang
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Erlang in section \"Initialize the client\""
# TODO(validate): erlang-server validator's gen_server harness is incompatible with the erlang-syntax-only scaffold's module shape. See _sdk-docs-port-notes.md.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
% This starts an instance with the default options
ldclient:start_instance("YOUR_SDK_KEY")

% You can also start a named instance
ldclient:start_instance("YOUR_SDK_KEY", your_instance)
```
