---
id: cpp-client-sdk/sdk-docs/features/privateattrs/config-c-sdk-v2
sdk: cpp-client-sdk
kind: reference
lang: c
description: Private attribute configuration for C client SDK v2.x.
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only-v2-c

---

```c
struct LDConfig *config = LDConfigNew("example-mobile-key");
// Mark all attributes private
LDConfigSetAllAttributesPrivate(config, true);
```
