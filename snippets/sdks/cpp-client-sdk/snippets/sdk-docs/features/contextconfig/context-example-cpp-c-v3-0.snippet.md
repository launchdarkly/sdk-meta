---
id: cpp-client-sdk/sdk-docs/features/contextconfig/context-example-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Context example for C++ (client-side) SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
/* Ensure array_builder.h is included to enable building array values for the context */
#include <launchdarkly/bindings/c/array_builder.h>

LDContextBuilder context_builder = LDContextBuilder_New();
LDContextBuilder_AddKind(context_builder, "user", "example-user-key");
LDContextBuilder_Attributes_Set(context_builder, "user", "firstName", LDValue_NewString("Sandy"));
LDContextBuilder_Attributes_Set(context_builder, "user", "lastName", LDValue_NewString("Smith"));

LDArrayBuilder group_builder = LDArrayBuilder_New();
LDArrayBuilder_Add(group_builder, LDValue_NewString("Acme"));
LDArrayBuilder_Add(group_builder, LDValue_NewString("Global Health Services"));

LDContextBuilder_Attributes_Set(context_builder, "user", "groups", LDArrayBuilder_Build(group_builder));

LDContext context = LDContextBuilder_Build(context_builder);
```
