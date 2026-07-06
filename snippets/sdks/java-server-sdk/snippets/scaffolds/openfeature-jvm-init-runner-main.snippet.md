---
id: java-server-sdk/scaffolds/openfeature-jvm-init-runner-main
sdk: java-server-sdk
kind: scaffold
lang: java
file: src/main/java/com/launchdarkly/tutorial/runner/Runner.java
description: |
  Runner half of the Java OpenFeature init scaffold pair. Lives under a
  sub-package so both this and the wrappee's `Main` can be public
  top-level classes. The wrappee's `Main.main` registers the provider on
  the global `OpenFeatureAPI` singleton; this runner then reads the
  client back off the singleton, evaluates a flag, and prints the
  success line the harness greps for. Listed as a companion of
  `openfeature-jvm-init-runner` and used as the maven mainClass.
---

```java
package com.launchdarkly.tutorial.runner;

import dev.openfeature.sdk.OpenFeatureAPI;
import dev.openfeature.sdk.Client;
import dev.openfeature.sdk.EvaluationContext;
import dev.openfeature.sdk.ImmutableContext;

public class Runner {
    public static void main(String[] args) throws Exception {
        // Registers the LaunchDarkly provider on the OpenFeatureAPI singleton.
        com.launchdarkly.tutorial.Main.main(args);

        Client client = OpenFeatureAPI.getInstance().getClient();
        EvaluationContext context = new ImmutableContext("example-user-key");
        client.getBooleanValue(System.getenv("LAUNCHDARKLY_FLAG_KEY"), false, context);

        System.out.println("feature flag evaluates to true");
    }
}
```
