---
id: erlang-server-sdk/sdk-docs/transport-layer-security-tls-erlang
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Erlang in section \"Transport Layer Security (TLS)\""
# Bucket C: erlang-server validator's gen_server harness is incompatible with the erlang-syntax-only scaffold's module shape. See _sdk-docs-port-notes.md.
---

```erlang
  ldclient:start_instance(SdkKey, #{
  http_options => #{
      tls_options => YourOptions
  }}),
```
