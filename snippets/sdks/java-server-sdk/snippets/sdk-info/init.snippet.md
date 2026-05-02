---
id: java-server-sdk/sdk-info/init
sdk: java-server-sdk
kind: init
lang: java
file: java-server-sdk/init.txt
description: Client initialization snippet for java-server-sdk.
validation:
  scaffold: java-server-sdk/scaffolds/init-runner
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```java
import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.server.*;

public class Main {
  public static void main(String[] args) {
    LDConfig config = new LDConfig.Builder().build();

    // This is your LaunchDarkly SDK key.
    // Never hardcode your SDK key in production.
    final LDClient client = new LDClient("YOUR_SDK_KEY", config);

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
