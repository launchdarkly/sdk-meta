---
id: cpp-client-sdk/sdk-docs/features/flagchanges/flag-changes-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Flag change subscription example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto listener = client.FlagNotifier().OnFlagChange("example-flag-key", [](auto event) {
  if (event->Deleted()) {
    std::cout << "The flag was deleted" << std::endl;
  } else {
    std::cout << "The flag was " << event->OldValue() << " and now it is " << event->NewValue() << std::endl;
  }
});

/* Then, you can disconnect the listener later */
listener->Disconnect();
```
