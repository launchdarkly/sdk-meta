---
id: cpp-client-sdk/sdk-docs/features/webproxy/web-proxy-cpp-native-v3-11
sdk: cpp-client-sdk
kind: reference
lang: cpp
description: Web proxy configuration for the C++ (client-side) SDK v3.11+ (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```cpp
// Requires an SDK build with -DLD_CURL_NETWORKING=ON.
auto config_builder = client_side::ConfigBuilder("example-mobile-key");
config_builder.HttpProperties().Proxy("https://web-proxy.domain.com:8080");
```
