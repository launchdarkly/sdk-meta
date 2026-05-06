---
id: cpp-server-sdk/sdk-docs/c-c---initialize-the-client-c-sdk-v3-0-c-binding
sdk: cpp-server-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Initialize the client\""
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```c
LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
     /* an error occurred, config is not valid */
}

LDServerSDK client = LDServerSDK_New(config);
```
