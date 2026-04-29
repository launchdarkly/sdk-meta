---
id: java-server-sdk/getting-started/mvn-generate
sdk: java-server-sdk
kind: bootstrap
lang: shell
description: Bootstrap a Maven project.
ld-application:
  slot: mvn-generate
---

Create a new project and accept the default options suggested by maven:

```shell
mvn archetype:generate -DgroupId=com.launchdarkly.tutorial -DartifactId=hello-java
```
