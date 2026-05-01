---
id: cpp-server-sdk/sdk-docs/c-c---initialize-the-client-c-sdk-v3-0-c-binding
sdk: cpp-server-sdk
kind: reference
lang: c
description: "C++ SDK v3.0 (C binding) in section \"Initialize the client\""
---

```c
LDClientConfigBuilder builder = LDClientConfigBuilder_New("YOUR_SDK_KEY");

LDClientConfig config;
LDStatus status = LDClientConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
     /* an error occurred, config is not valid */
}

LDClientSDK client = LDClientSDK_New(config);
```
