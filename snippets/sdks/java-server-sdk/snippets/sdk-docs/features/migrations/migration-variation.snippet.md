---
id: java-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: java-server-sdk
kind: reference
lang: java
description: Migration stage evaluation (migrationVariation) for Java SDK v7.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
LDContext context = LDContext.builder("example-context-key")
  .build();

MigrationVariation migrationVariation = client.migrationVariation("example-migration-flag-key", context, MigrationStage.OFF);
```
