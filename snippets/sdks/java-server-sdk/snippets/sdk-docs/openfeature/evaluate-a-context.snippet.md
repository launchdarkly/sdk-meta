---
id: java-server-sdk/sdk-docs/openfeature/evaluate-a-context
sdk: java-server-sdk
kind: reference
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: "Java OpenFeature provider in section \"Evaluate a context\""
validation:
  scaffold: java-server-sdk/scaffolds/openfeature-jvm-runner
---

```java
boolean value = client.getBooleanValue("example-flag-key", false, context);
```
