---
id: erlang-server-sdk/sdk-docs/get-started-erlang-2
sdk: erlang-server-sdk
kind: reference
lang: erlang
description: "Erlang in section \"Get started\""
# TODO(snippet-bug): body is a partial tuple (`{applications, [...]},`
# — note the trailing comma) extracted from a `.app.src` file, not
# a complete Erlang module. erlc can't compile it as a standalone
# `.erl` file. Fix in the snippet-bugs PR: either re-tag and route
# through a `.app.src` parser, or restructure the doc to wrap the
# fragment in the surrounding `{application, App, [...]}` tuple.
---

```erlang
{applications,
  [kernel,
  stdlib,
  ldclient
]},
```
