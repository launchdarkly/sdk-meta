---
id: java-server-sdk/getting-started/run
sdk: java-server-sdk
kind: run
lang: shell
description: Build the assembly jar and run with the SDK key in the environment.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key embedded in the rendered Run command.
ld-application:
  slot: run
---

Build and run:

```shell
mvn clean compile assembly:single && LAUNCHDARKLY_SDK_KEY={{ apiKey }} java -jar target/hello-java-1.0-SNAPSHOT-jar-with-dependencies.jar
```
