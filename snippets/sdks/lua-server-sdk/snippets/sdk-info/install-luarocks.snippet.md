---
id: lua-server-sdk/sdk-info/install-luarocks
sdk: lua-server-sdk
kind: install
lang: shell
file: lua-server-sdk/install-luarocks.txt
description: |
  Install fragment for lua-server-sdk (luarocks). Unvalidated: the
  shell-install harness has no luarocks toolchain, and the rock links
  against the LaunchDarkly C++ server SDK, which must be built and
  installed first (LD_DIR points at it).
---

```shell
# The Lua SDK wraps the LaunchDarkly C++ server-side SDK; install that
# dynamic library first, then point LD_DIR at it.
luarocks install launchdarkly-server-sdk LD_DIR=/path/to/cpp-sdk/install
```
