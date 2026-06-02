---
id: erlang-server-sdk/sdk-docs/get-started-elixir
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Elixir in section \"Get started\""
# TODO(validate): erlang-server validator's gen_server harness is incompatible with the erlang-syntax-only scaffold's module shape. See _sdk-docs-port-notes.md.
validation:
  scaffold: erlang-server-sdk/scaffolds/erlang-syntax-only
---

```erlang
defp deps do
  [
    {:ldclient, "~> 3.0.0", hex: :launchdarkly_server_sdk}
  ]
end
```
