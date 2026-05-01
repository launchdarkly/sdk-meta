---
id: cpp-client-sdk/sdk-docs/initialize-the-client-c-sdk-v3-0-c-binding
sdk: cpp-client-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Initialize the client\""
---

```c
LDClientConfigBuilder builder = LDClientConfigBuilder_New("example-mobile-key");

LDClientConfig config;
LDStatus status = LDClientConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
     /* an error occurred, config is not valid */
}

LDContextBuilder context_builder = LDContextBuilder_New();
LDContextBuilder_AddKind(context_builder, "user", "example-user-key");

LDContext context = LDContextBuilder_Build(context_builder);
```
