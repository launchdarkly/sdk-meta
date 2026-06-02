---
id: erlang-server-sdk/sdk-docs/get-started-erlang
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Erlang in section \"Get started\""
# TODO(snippet-bug): body is rebar.config syntax (a `{deps, [...]}.`
# top-level tuple), not an Erlang module. erlc compiles `.erl`
# files containing `-module(...)` + function declarations; this
# fragment would belong in `rebar.config`. Fix in the snippet-bugs
# PR: re-tag and route through a rebar.config parser, or rewrite
# the body as a module that documents the dependency in comments.
---

```erlang
{deps, [
  {ldclient, "3.0.0", {pkg, launchdarkly_server_sdk}}
]}.
```
