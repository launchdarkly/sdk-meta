---
id: erlang-server-sdk/getting-started/sup-childspecs
sdk: erlang-server-sdk
kind: manifest-fragment
lang: erlang
description: ChildSpecs replacement to drop into src/hello_erlang_sup.erl.
ld-application:
  slot: sup-childspecs
---

Replace the `ChildSpecs` variable in `src/hello_erlang_sup.erl` with the following:

```erlang
[{console,
            {hello_erlang_server, start_link, []},
            permanent, 5000, worker, [hello_erlang_server]}]
```
