---
id: cpp-server-sdk/sdk-docs/migration-2-to-3-introducing-the-c-server-side-sdk-as-a-replacement-for-the-c-server-side-sdk-c-sdk-v3-0-native
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Introducing the C++ (server-side) SDK as a replacement for the C (server-side) SDK\""
---

```cpp
if (auto const status  = start_result.wait_for(maxwait); status == std::future_status::ready) {
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
