---
id: android-client-sdk/sdk-docs/features/monitoring/connection-info-kotlin
sdk: android-client-sdk
kind: reference
lang: kotlin
description: Connection information retrieval for Android (Kotlin).
validation:
  scaffold: android-client-sdk/scaffolds/kotlin-syntax-only

---

```kotlin
val client: LDClient = LDClient.get()
val connectionInfo: ConnectionInformation = client.connectionInformation
// One of the seven modes described above
val connectionMode: ConnectionInformation.ConnectionMode =
    connectionInfo.connectionMode
// Most recent successful flag cache update in millis from the epoch
// Or null if flags have never been retrieved
val lastSuccess: Long? = connectionInfo.lastSuccessfulConnection
// Most recent unsuccessful flag cache update attempt in millis from the epoch
// Or null if flag update has never been attempted
val lastError: Long? = connectionInfo.lastFailedConnection
// Most recent failure or null
val ldFailure: LDFailure? = connectionInfo.lastFailure
```
