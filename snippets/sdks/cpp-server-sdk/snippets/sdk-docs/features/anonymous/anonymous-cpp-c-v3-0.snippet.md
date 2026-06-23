---
id: cpp-server-sdk/sdk-docs/features/anonymous/anonymous-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Anonymous context example for C++ (server-side), C binding.
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```c
LDContextBuilder builder = LDContextBuilder_New();

LDContextBuilder_AddKind(builder, "user", "example-user-key");
LDContextBuilder_Attributes_SetAnonymous(builder, "user", true);

LDContext context = LDContextBuilder_Build(builder);
```
