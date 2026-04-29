---
id: erlang-server-sdk/getting-started/app-src
sdk: erlang-server-sdk
kind: manifest-fragment
lang: erlang
description: applications block to add to src/hello_erlang.app.src.
ld-application:
  slot: app-src
---

Edit `src/hello_erlang.app.src` to import LaunchDarkly:

```erlang
{applications,
	[kernel,
	stdlib,
	ldclient
]},
```
