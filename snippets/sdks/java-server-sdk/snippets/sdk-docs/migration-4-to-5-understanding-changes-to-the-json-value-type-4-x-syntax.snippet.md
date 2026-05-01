---
id: java-server-sdk/sdk-docs/migration-4-to-5-understanding-changes-to-the-json-value-type-4-x-syntax
sdk: java-server-sdk
kind: reference
lang: java
description: "4.x syntax in section \"Understanding changes to the JSON value type\""
---

```java
// 4.x way: set user's "groups" to ["cats", "dogs"]
JsonArray groups = new JsonArray();
groups.add("cats");
groups.add("dogs");
LDUser user = new LDUser.Builder("key")
  .custom("group", groups)
  .build();

// The following shortcut method was exactly equivalent:
LDUser user = new LDUser.Builder("key")
  .customString("group", Arrays.asList("cats", "dogs"))
  .build();
```
