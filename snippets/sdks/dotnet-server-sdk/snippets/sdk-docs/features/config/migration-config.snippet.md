---
id: dotnet-server-sdk/sdk-docs/features/config/migration-config
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: Migration configuration example for the .NET (server-side) SDK v8 — read/write methods, execution order, latency/error tracking.
validation:
  scaffold: dotnet-server-sdk/scaffolds/csharp-syntax-only

---

```csharp

// define how to compare the two read values
bool Checker(string a, string b) => a.Equals(b);

var migration = new MigrationBuilder<string, string, string, string>(_client)
    .Read(
        (payload) => MigrationMethod.Success("read old"),
        (payload) => MigrationMethod.Success("read new"),
        Checker)
    .Write(
        (payload) => MigrationMethod.Success("write old"),
        (payload) => MigrationMethod.Success("write new"))
    .ReadExecution(MigrationExecution.Parallel()) // or MigrationExecution.Serial(MigrationSerialOrder.Fixed)
    .TrackErrors(true)    // true by default
    .TrackLatency(true)   // true by default
    .Build();

```
