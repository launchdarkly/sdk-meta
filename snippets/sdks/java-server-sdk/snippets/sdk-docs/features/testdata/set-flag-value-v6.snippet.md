---
id: java-server-sdk/sdk-docs/features/testdata/set-flag-value-v6
sdk: java-server-sdk
kind: reference
lang: java
description: Setting a test data flag to a specific value for Java SDK v6.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
td.update(td.flag("example-flag-key").variationForAll(false));
```
