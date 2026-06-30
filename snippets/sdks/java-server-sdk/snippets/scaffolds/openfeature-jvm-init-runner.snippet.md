---
id: java-server-sdk/scaffolds/openfeature-jvm-init-runner
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/tutorial/Main.java
description: |
  Compiles and runs an OpenFeature "initialize the provider" fragment
  end-to-end against a real LaunchDarkly environment. The fragment is a
  full `public class Main` whose `main` registers a LaunchDarkly provider
  with the global `OpenFeatureAPI` singleton; this scaffold supplies the
  `import` statements the page's "import the namespaces" fragment
  documents so the body compiles. The runner companion is the maven
  entrypoint: it invokes the body's `Main.main`, then reads the now-
  registered client back off the `OpenFeatureAPI` singleton, evaluates a
  flag, and prints the success line. The fragment's `YOUR_SDK_KEY`
  literal is replaced with the real key via `validation.placeholders`.
inputs:
  body:
    type: string
    description: The wrappee init fragment; a full `public class Main` that registers the provider.
validation:
  runtime: jvm
  entrypoint: src/main/java/com/launchdarkly/tutorial/runner/Runner.java
  companions:
    - java-server-sdk/scaffolds/openfeature-jvm-init-runner-main
---

```java
package com.launchdarkly.tutorial;

import dev.openfeature.sdk.OpenFeatureAPI;
import dev.openfeature.sdk.Client;
import com.launchdarkly.sdk.server.LDClient;
import com.launchdarkly.openfeature.serverprovider.Provider;

{{ body }}
```
