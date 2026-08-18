---
id: cpp-client-sdk/sdk-docs/features/webproxy/web-proxy-auth-cpp-c-v3-11
sdk: cpp-client-sdk
kind: reference
lang: c
description: Web proxy configuration with authentication for the C++ (client-side) SDK v3.11+ (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
/* Requires an SDK build with -DLD_CURL_NETWORKING=ON. */
LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_HttpProperties_Proxy(
    builder, "https://username:password@web-proxy.domain.com:8080");
```
