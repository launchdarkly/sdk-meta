---
id: java-server-sdk/sdk-docs/migration-6-to-7-configuring-the-migration-java-sdk-v7-0
sdk: java-server-sdk
kind: reference
lang: java
description: "Java SDK v7.0 in section \"Configuring the migration\""
---

```java

LDClient client = new LDClient("YOUR_SDK_KEY");

MigrationBuilder<String, String, String, String> migrationBuilder = new MigrationBuilder<>(client)
  .read(
    (payload) -> MigrationMethodResult.Success("read old"),
    (payload) -> MigrationMethodResult.Success("read new"),
    (a, b) -> a.equals(b)
  )
  .readExecution(MigrationExecution.Serial(MigrationSerialOrder.RANDOM)) // default is .Parallel
  .write(
    (payload) -> MigrationMethodResult.Success("write old"),
    (payload) -> MigrationMethodResult.Success("write new")
  )
  .trackLatency(true)
  .trackErrors(true)
  .build();

Migration<String, String, String, String> migration = migrationBuilder.build();

```
