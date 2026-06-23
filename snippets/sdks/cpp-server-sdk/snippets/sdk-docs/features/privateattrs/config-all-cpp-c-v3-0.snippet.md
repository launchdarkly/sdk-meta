---
id: cpp-server-sdk/sdk-docs/features/privateattrs/config-all-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Marking all attributes private for C++ server SDK v3.0 (C binding).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only

---

```c
/* sets all attributes private */
LDServerConfigBuilder config_builder = LDServerConfigBuilder_New("example-mobile-key");
LDServerConfigBuilder_Events_AllAttributesPrivate(config_builder, true);

LDServerConfig config;
LDStatus config_status = LDServerConfigBuilder_Build(config_builder, &config);

/* sets "email" and "address" private */
config_builder = LDServerConfigBuilder_New("example-mobile-key");
LDServerConfigBuilder_Events_PrivateAttribute(config_builder, "email");
LDServerConfigBuilder_Events_PrivateAttribute(config_builder, "address");

config_status = LDServerConfigBuilder_Build(config_builder, &config);
```
