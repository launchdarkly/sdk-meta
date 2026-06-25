---
id: cpp-client-sdk/sdk-docs/features/webproxy/web-proxy-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: Web proxy configuration for the C client SDK v2.x (native).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c
---

```c
struct LDConfig *config = LDConfigNew("example-mobile-key");
LDConfigSetProxyURI(config, "https://web-proxy.domain.com:8080");
```
