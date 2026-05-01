---
id: android-client-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-anonymous-users-4-0-syntax-configuring-the-sdk-to-generate-keys
sdk: android-client-sdk
kind: reference
lang: java
description: "4.0 syntax, configuring the SDK to generate keys in section \"Understanding changes to anonymous users\""
---

```java
LDConfig config = new LDConfig.Builder()
    .generateAnonymousKeys(true)
    .build();
```
