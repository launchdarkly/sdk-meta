---
id: lua-server-sdk/getting-started/luarocks-install
sdk: lua-server-sdk
kind: install
lang: bash
description: Install the Lua server SDK via luarocks.
ld-application:
  slot: luarocks-install
---

Download the Lua Server SDK and build it with `luarocks` (replace `LD_DIR` with the path to the C++ SDK's shared libraries as necessary):

```bash
luarocks install launchdarkly-server-sdk LD_DIR="$(pwd)/cpp-sdks/build/install"
```
