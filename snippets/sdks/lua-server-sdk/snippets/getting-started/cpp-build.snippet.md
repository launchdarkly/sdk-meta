---
id: lua-server-sdk/getting-started/cpp-build
sdk: lua-server-sdk
kind: install
lang: bash
description: Compile and install the underlying C++ Server SDK from source.
ld-application:
  slot: cpp-build
---

If the C++ Server SDK is already installed or you already obtained release artifacts from LaunchDarkly, skip this step.

Otherwise, compile and install the C++ Server SDK:

```bash

git clone https://github.com/launchdarkly/cpp-sdks.git && cd cpp-sdks
mkdir build && cd build
cmake -G Ninja -D BUILD_TESTING=OFF \
               -D CMAKE_BUILD_TYPE=Release \
               -D LD_BUILD_SHARED_LIBS=On \
               -D CMAKE_INSTALL_PREFIX=./install ..
cmake --build . --target launchdarkly-cpp-server
cmake --install .
cd ../../
        
```
