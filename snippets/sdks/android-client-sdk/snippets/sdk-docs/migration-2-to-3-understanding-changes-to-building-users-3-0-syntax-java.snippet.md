---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-building-users-3-0-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "3.0 syntax (Java) in section \"Understanding changes to building users\""
---

```java
LDUser user = new LDUser.Builder("userKey")
    .custom("properties", LDValue.buildArray().add("new").add("priority").build())
    .privateCustom("counts", LDValue.buildArray().add(3).add(5).build())
    .build();
```
