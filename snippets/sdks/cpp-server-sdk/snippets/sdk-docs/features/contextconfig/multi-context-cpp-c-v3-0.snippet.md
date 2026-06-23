---
id: cpp-server-sdk/sdk-docs/features/contextconfig/multi-context-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: cpp
description: Multi-context example for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```cpp
LDContextBuilder context_builder = LDContextBuilder_New();
LDContextBuilder_AddKind(context_builder, "user", "example-user-key");
LDContextBuilder_Attributes_SetName(context_builder, "user", "Sandy");
LDContextBuilder_AddKind(context_builder, "organization", "example-organization-key");
LDContextBuilder_Attributes_SetName(context_builder, "organization", "Global Health Services");
LDContext context = LDContextBuilder_Build(context_builder);
```
