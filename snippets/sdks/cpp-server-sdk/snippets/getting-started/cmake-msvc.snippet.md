---
id: cpp-server-sdk/getting-started/cmake-msvc
sdk: cpp-server-sdk
kind: bootstrap
lang: bash
description: Configure with Visual Studio 2022.
ld-application:
  slot: cmake-msvc
---

```bash
cmake -G"Visual Studio 17 2022" -DBUILD_TESTING=OFF ..
```
