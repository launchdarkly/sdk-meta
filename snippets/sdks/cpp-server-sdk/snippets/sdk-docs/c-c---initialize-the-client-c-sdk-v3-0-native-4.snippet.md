---
id: cpp-server-sdk/sdk-docs/c-c---initialize-the-client-c-sdk-v3-0-native-4
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Initialize the client\""
---

```cpp
server_side::Client client(*config);

auto start_result = client.StartAsync();
if (auto const status = start_result.wait_for(maxwait); status == std::future_status::ready) {
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
