---
id: cpp-client-sdk/sdk-docs/features/identify/identify-result-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Examining the identify result for the C++ client SDK v3.0 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
auto identify_result = client.IdentifyAsync(updated_context);
auto status = identify_result.wait_for(maxwait);

if (status == std::future_status::ready) {
  /* The client's attempt to identify succeeded or failed in the specified amount of time. */
  if (identify_result.get()) {
    /* Identification succeeded */
  } else {
    /* Identification failed */
  }
} else {
  /* The specified timeout was reached, but the client is still identifying. */
}
```
