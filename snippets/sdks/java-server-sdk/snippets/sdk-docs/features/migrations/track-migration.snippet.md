---
id: java-server-sdk/sdk-docs/features/migrations/track-migration
sdk: java-server-sdk
kind: reference
lang: java
description: Migration metrics tracking (trackMigration) for Java SDK v7.0.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
MigrationOpTracker tracker = migrationVariation.getTracker();

client.trackMigration(tracker);
```
