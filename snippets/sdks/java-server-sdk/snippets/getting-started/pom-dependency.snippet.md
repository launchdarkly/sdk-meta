---
id: java-server-sdk/getting-started/pom-dependency
sdk: java-server-sdk
kind: manifest-fragment
lang: xml
description: pom.xml `<dependencies>` entry for the Java server SDK.
inputs:
  version:
    type: string
    description: SDK version. Gonfalon fetches the latest from Maven Central asynchronously; renders as empty during the fetch (rather than the prior stale '5.0.0' fallback).
    runtime-default: ""
ld-application:
  slot: pom-dependency
---

Add the SDK to your project in your `pom.xml <dependencies>` section:

```xml
<dependency>
  <groupId>com.launchdarkly</groupId>
  <artifactId>launchdarkly-java-server-sdk</artifactId>
  <version>{{ version }}</version>
</dependency>
```
