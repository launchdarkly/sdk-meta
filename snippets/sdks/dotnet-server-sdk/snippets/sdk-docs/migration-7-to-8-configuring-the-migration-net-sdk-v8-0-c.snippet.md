---
id: dotnet-server-sdk/sdk-docs/migration-7-to-8-configuring-the-migration-net-sdk-v8-0-c
sdk: dotnet-server-sdk
kind: reference
lang: csharp
description: ".NET SDK v8.0 (C#) in section \"Configuring the migration\""
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
