---
id: cpp-client-sdk/sdk-docs/migration-2-to-3-introducing-the-c-client-side-sdk-as-a-replacement-for-the-c-client-side-sdk-c-sdk-v3-0-native
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Introducing the C++ (client-side) SDK as a replacement for the C (client-side) SDK\""
---

```cpp
auto start_result = client.StartAsync();
auto status = start_result.wait_for(maxwait);
if (status == std::future_status::ready) {
    /* The client's attempt to initialize succeeded or failed in the specified amount of time. */
    if (start_result.get()) {
        /* Initialization succeeded. */
    } else {
        /* Initialization failed. */
    }
} else {
    /* The specified timeout was reached, but the client is still initializing. */
}
```
