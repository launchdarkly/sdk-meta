---
id: cpp-client-sdk/sdk-docs/features/privateattrs/config-cpp-c-v3-0
sdk: cpp-client-sdk
kind: reference
lang: c
description: Private attribute configuration for C++ client SDK v3.0 (C binding).
validation:
  scaffold: cpp-client-sdk/scaffolds/cpp-client-syntax-only

---

```c
/* sets all attributes private */
LDClientConfigBuilder config_builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_Events_AllAttributesPrivate(config_builder, true);

LDClientConfig config;
LDStatus config_status = LDClientConfigBuilder_Build(config_builder, &config);

/* sets "email" and "address" private */
config_builder = LDClientConfigBuilder_New("example-mobile-key");
LDClientConfigBuilder_Events_PrivateAttribute(config_builder, "email");
LDClientConfigBuilder_Events_PrivateAttribute(config_builder, "address");

config_status = LDClientConfigBuilder_Build(config_builder, &config);
```
