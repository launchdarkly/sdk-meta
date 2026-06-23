---
id: apex-server-sdk/sdk-docs/features/privateattrs/user
sdk: apex-server-sdk
kind: reference
lang: java
description: Marking user attributes private with the user builder for Apex.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only

---

```java
Set<String> privateAttributes = new Set<String>();
privateAttributes.add('firstName');

LDUser user = new LDUser.Builder('example-user-key')
    .setFirstName('alice')
    .setPrivateAttributeNames(privateAttributes)
    .build();
```
