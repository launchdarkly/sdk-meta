---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v3-0-native-2
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Initialize the client\""
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only
---

```cpp
client_side::Client client(config, context);
client.StartAsync().wait_for(std::chrono::seconds(10));
```
