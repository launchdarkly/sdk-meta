---
id: cpp-server-sdk/sdk-docs/c-c---evaluate-a-context-c-sdk-v3-0-c-binding
sdk: cpp-server-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Evaluate a context\""
---

```c
LDContextBuilder context_builder = LDContextBuilder_New();
LDContextBuilder_AddKind(context_builder, "user", "example-user-key");
LDContextBuilder_Attributes_SetName(context_builder, "user", "Sandy");
LDContext context = LDContextBuilder_Build(context_builder);

bool show_feature = LDServerSDK_BoolVariation(client, context, "example-flag-key", false);

if (show_feature) {
    // application code to show the feature
} else {
    // the code to run if the feature is off
}
```
