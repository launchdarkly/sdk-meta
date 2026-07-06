---
id: java-server-sdk/sdk-docs/openfeature/construct-a-context-organization
sdk: java-server-sdk
kind: reference
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: "Java OpenFeature provider in section \"Construct a context\" (organization)"
validation:
  scaffold: java-server-sdk/scaffolds/openfeature-jvm-context-runner
---

```java
EvaluationContext context = new ImmutableContext("org-key", new HashMap<String, Value>(){{
    put("kind", new Value("organization"));
}});
```
