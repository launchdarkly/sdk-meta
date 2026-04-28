---
id: cpp-client-sdk/getting-started/cmake-ninja
sdk: cpp-client-sdk
kind: bootstrap
lang: bash
description: Configure with the Ninja generator.
ld-application:
  slot: cmake-ninja
---

```bash
cmake -G"Ninja" -DBUILD_TESTING=OFF ..
```
