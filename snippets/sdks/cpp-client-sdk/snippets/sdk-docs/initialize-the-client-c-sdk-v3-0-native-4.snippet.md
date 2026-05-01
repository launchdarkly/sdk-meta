---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v3-0-native-4
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Initialize the client\""
---

```cpp
client_side::Client client(config, context);

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
