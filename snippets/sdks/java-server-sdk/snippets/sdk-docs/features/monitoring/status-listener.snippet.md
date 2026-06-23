---
id: java-server-sdk/sdk-docs/features/monitoring/status-listener
sdk: java-server-sdk
kind: reference
lang: java
description: Data source status change listener for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only

---

```java
client.getDataSourceStatusProvider().addStatusListener(
  status -> {
    System.out.println("new status is: " + status);
  }
);
```
