---
id: java-server-sdk/sdk-docs/features/migrations/read-write
sdk: java-server-sdk
kind: reference
lang: java
description: Migration read and write example for Java SDK v7.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
LDContext context = LDContext.builder("example-context-key")
  .build();

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
MigrationStage defaultStage = MigrationStage.OFF;

Migration.MigrationResult<String> readResult = migration.read("example-migration-flag-key", context, defaultStage);

Migration.MigrationWriteResult<String> writeResult = migration.write("example-migration-flag-key", context, defaultStage);
```
