---
id: erlang-server-sdk/sdk-info/install-rebar
sdk: erlang-server-sdk
kind: install
lang: erlang
file: erlang-server-sdk/install-rebar.txt
description: |
  Install fragment for erlang-server-sdk (rebar.config deps). A
  declarative manifest fragment, not runnable Erlang — same unvalidated
  disposition as the haskell cabal/stack and java gradle/xml fragments.
  The hex package is launchdarkly_server_sdk, aliased locally as
  ldclient.
---

```erlang
{deps, [
  {ldclient, "3.11.2", {pkg, launchdarkly_server_sdk}}
]}.
```
