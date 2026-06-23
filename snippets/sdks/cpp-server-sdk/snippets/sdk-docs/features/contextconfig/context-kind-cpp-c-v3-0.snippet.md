---
id: cpp-server-sdk/sdk-docs/features/contextconfig/context-kind-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Context with a non-user kind for C++ (server-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
LDContextBuilder context_builder = LDContextBuilder_New();
LDContextBuilder_AddKind(context_builder, "organization", "example-organization-key");

LDContext context = LDContextBuilder_Build(context_builder);
```
