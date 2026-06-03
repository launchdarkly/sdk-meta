---
id: erlang-server-sdk/sdk-docs/transport-layer-security-tls-erlang
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Erlang in section \"Transport Layer Security (TLS)\""
# TODO(snippet-bug): body ends with a trailing `,` — it's a fragment
# meant to be spliced inside a larger function body. The
# erlang-syntax-only scaffold wraps `{{ body }}` with `wrappee_() ->`
# / `.`, which yields `..., .` (illegal). Fix in the snippet-bugs
# PR: end with `.` or wrap as a complete expression statement so
# the scaffold's terminator is valid.
---

```erlang
  ldclient:start_instance(SdkKey, #{
  http_options => #{
      tls_options => YourOptions
  }}),
```
