---
id: cpp-client-sdk/sdk-docs/features/contextconfig/context-kind-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Context with a non-user kind for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
LDContextBuilder context_builder = LDContextBuilder_New();
LDContextBuilder_AddKind(context_builder, "organization", "example-organization-key");

LDContext context = LDContextBuilder_Build(context_builder);
```
