---
id: apex-server-sdk/sdk-docs/features/contextconfig/user-example
sdk: apex-server-sdk
kind: reference
lang: java
description: User example for Apex SDK.
validation:
  scaffold: apex-server-sdk/scaffolds/apex-syntax-only

---

```java
LDUser user = new LDUser.Builder('example-user-key')
    .setFirstName('Sandy')
    .setLastName('Smith')
    .setEmail('sandy@example.com')
    .setCustom(new LDValueObject.Builder()
        .set('groups', new LDValueArray.Builder()
            .add(LDValue.of('Acme'))
            .add(LDValue.of('Global Health Services'))
            .build()
        )
        .build()
    )
    .build();
```
