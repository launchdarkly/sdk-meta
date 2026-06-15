---
id: cpp-client-sdk/sdk-docs/features/allflags/allflags-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: All flags example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
for (auto [flag_key, flag_value] : client.AllFlags()) {
    std::cout << flag_key << ": " << flag_value << std::endl;
}
```
