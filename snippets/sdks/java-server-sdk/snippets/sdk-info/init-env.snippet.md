---
id: java-server-sdk/sdk-info/init-env
sdk: java-server-sdk
kind: init
lang: java
file: java-server-sdk/init-env.txt
description: Client initialization snippet for java-server-sdk reading the SDK key from an environment variable.
validation:
  scaffold: java-server-sdk/scaffolds/init-runner
---

```java
import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.server.*;

public class Main {
  public static void main(String[] args) {
    LDConfig config = new LDConfig.Builder().build();

    // Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
    // environment variable, so one build can run in every environment.
    final LDClient client = new LDClient(System.getenv("LAUNCHDARKLY_SDK_KEY"), config);

    if (client.isInitialized()) {
      // For onboarding purposes only we flush events as soon as
      // possible so we quickly detect your connection.
      // You don't have to do this in practice because events are automatically flushed.
      client.flush();
      System.out.println("SDK successfully initialized!");
    }
  }
}
```
