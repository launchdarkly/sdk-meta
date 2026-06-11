---
id: cpp-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Flag evaluation reason example for C++ (client-side) SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
auto detail = client.BoolVariationDetail("example-flag-key", false);
if (detail.Value()) {
  std::cout << "Value was true!" << std::endl;
} else {
  // it was false, let's find out why.
  if (auto reason = detail.Reason(); reason.has_value()) {
    // reason might not be present, so we have to check
    std::cout << "Value was false because of " << reason.value() << std::endl;
  } else {
    std::cout << "No reason provided to explain why flag was false!" << std::endl;
  }
}
```
