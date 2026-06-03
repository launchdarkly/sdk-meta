---
id: apex-server-sdk/sdk-docs/apex-initialize-the-client-custom-config
sdk: apex-server-sdk
kind: reference
lang: java
description: "Custom Config in section \"Initialize the client\""
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only
---

```java
LDConfig config = new LDConfig.Builder()
    .setAllAttributesPrivate(true)
    .build();
LDClient client = new LDClient(config);
```
