---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-the-json-value-type-5-0-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "5.0 syntax in section \"Understanding changes to the JSON value type\""
---

```java
// 5.0 way: set user's "groups" to ["cats", "dogs"]
LDValue groups = LDValue.buildArray().add("cats").add("dogs").build();
LDUser user = new LDUser.Builder("key")
  .custom("groups", groups)
  .build();
```
