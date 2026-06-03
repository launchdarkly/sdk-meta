---
id: java-server-sdk/sdk-docs/features/config/migration-config
sdk: java-server-sdk
kind: reference
lang: java
description: Migration configuration example for the Java SDK v7 — read/write methods, execution order, latency/error tracking.

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
  .trackErrors(true);

Migration<String, String, String, String> migration = migrationBuilder.build();

```
