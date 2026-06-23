---
id: cpp-client-sdk/sdk-docs/features/identify/identify-update-context-cpp-native-v3-0
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Updating the initial context before an identify call, for the C++ client SDK v3 (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
/* before end user logs in */
auto context = ContextBuilder()
  .Kind("device", "example-device-key")
  .Build();

/* after end user logs in */
auto updated_context = ContextBuilder(context)
  .Kind("user", "example-user-key")
  .Name("Sandy")
  .Kind("organization", "example-organization-key")
  .Name("Global Health Services")
  .Build();
```
