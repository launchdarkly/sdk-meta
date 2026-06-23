---
id: java-server-sdk/sdk-docs/features/monitoring/check-status
sdk: java-server-sdk
kind: reference
lang: java
description: Data source status check for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
DataSourceStatusProvider.Status dataSourceStatus =
  client.getDataSourceStatusProvider().getStatus();
```
