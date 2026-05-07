---
id: cpp-server-sdk/sdk-docs/c-c---initialize-the-client-c-sdk-v3-0-native
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: "C++ SDK v3.0 (native) in section \"Initialize the client\""
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```cpp
auto config_builder = server_side::ConfigBuilder("YOUR_SDK_KEY");
auto config = config_builder.Build();
if (!config) {
   /* an error occurred, config is not valid */
}

server_side::Client client(*config);
```
