---
id: android-client-sdk/sdk-docs/features/monitoring/connection-info-java
sdk: android-client-sdk
kind: reference
lang: java
description: Connection information retrieval for Android (Java).
validation:
  scaffold: android-client-sdk/scaffolds/java-syntax-only

---

```java
LDClient client = LDClient.get();
ConnectionInformation connectionInfo = client.getConnectionInformation();
// One of the seven modes described above
ConnectionInformation.ConnectionMode connectionMode =
  connectionInfo.getConnectionMode();
// Most recent successful flag cache update in millis from the epoch
// Or null if flags have never been retrieved
Long lastSuccess = connectionInfo.getLastSuccessfulConnection();
// Most recent unsuccessful flag cache update attempt in millis from the epoch
// Or null if flag update has never been attempted
Long lastError = connectionInfo.getLastFailedConnection();
// Most recent failure or null
LDFailure ldFailure = connectionInfo.getLastFailure();
```
