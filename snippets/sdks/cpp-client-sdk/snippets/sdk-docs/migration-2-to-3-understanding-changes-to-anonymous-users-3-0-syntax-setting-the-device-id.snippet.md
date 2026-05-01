---
id: cpp-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-anonymous-users-3-0-syntax-setting-the-device-id
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "3.0 syntax, setting the device ID in section \"Understanding changes to anonymous users\""
---

```cpp
auto context = ContextBuilder()
  .Kind("device", "example-device-key")
  .Set("deviceID", "example-device-ID")
  .Build();
```
