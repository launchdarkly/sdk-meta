---
id: android-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-building-users-2-x-syntax-java
sdk: android-client-sdk
kind: reference
lang: java
description: "2.x syntax (Java) in section \"Understanding changes to building users\""
---

```java
LDUser user = new LDUser.Builder("userKey")
    .customString("properties", Arrays.asList("new", "priority"))
    .privateCustomNumber("counts", Arrays.asList(3, 5))
    .build();
```
