---
id: erlang-server-sdk/getting-started/rebar-config
sdk: erlang-server-sdk
kind: manifest-fragment
lang: erlang
description: rebar.config dependency entry for the LaunchDarkly Erlang server SDK.
inputs:
  version:
    type: string
    description: SDK version. Gonfalon fetches the latest from Hex asynchronously.
    runtime-default: ""
ld-application:
  slot: rebar-config
---

Next, add the SDK package to your list of dependencies in `rebar.config`:

```erlang
{ldclient, "{{ version }}", {pkg, launchdarkly_server_sdk}}
```
