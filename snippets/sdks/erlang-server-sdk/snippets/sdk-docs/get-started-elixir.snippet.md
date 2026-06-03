---
id: erlang-server-sdk/sdk-docs/get-started-elixir
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Elixir in section \"Get started\""
# TODO(snippet-bug): body is Elixir (`defp deps do ... end`), not
# Erlang — the source MDX tags this in an `erlang` code block but
# the actual content is Mix dependency-list syntax. erlc rejects it
# at parse. Fix in the snippet-bugs PR: re-tag as `elixir` and route
# through an elixir-specific scaffold (none exists yet), or rewrite
# the body in Erlang's `rebar.config` syntax.
---

```erlang
defp deps do
  [
    {:ldclient, "~> 3.0.0", hex: :launchdarkly_server_sdk}
  ]
end
```
