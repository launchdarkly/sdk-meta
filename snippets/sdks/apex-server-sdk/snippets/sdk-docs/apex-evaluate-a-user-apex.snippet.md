---
id: apex-server-sdk/sdk-docs/apex-evaluate-a-user-apex
sdk: apex-server-sdk
kind: reference
lang: java
description: "Apex in section \"Evaluate a user\""
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only
---

```java
LDUser user = new LDUser.Builder('example-user-key')
    .setName('Sandy')
    .build();

Boolean value = client.boolVariation(user, flagKey, false);
if (value) {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
