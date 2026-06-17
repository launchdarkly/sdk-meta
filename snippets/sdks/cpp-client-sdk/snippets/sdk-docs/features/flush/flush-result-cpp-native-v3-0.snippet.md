---
id: cpp-client-sdk/sdk-docs/features/flush/flush-result-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Flush result example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
/* FlushAsync returns immediately; event delivery proceeds in the
   background. The call does not report a result. */
client.FlushAsync();
```
