---
id: java-server-sdk/sdk-docs/migration-6-to-7-reading-and-writing-during-the-migration-java-sdk-v7-0
sdk: java-server-sdk
kind: reference
lang: java
description: "Java SDK v7.0 in section \"Reading and writing during the migration\""
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
