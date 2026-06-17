---
id: cpp-server-sdk/sdk-docs/features/allflags/allflags-cpp-native-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: All flags example for C++ (server-side) SDK v3.0 (native).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
auto const all_flags = client.AllFlagsState(context);
if (all_flags.Valid()) {
   for (auto const& [flag_key, flag_value] : all_flags.Values()) {
       std::cout << flag_key << ": " << flag_value << std::endl;
   }
}  else {
   /* error evaluating all flags! */
}
```
