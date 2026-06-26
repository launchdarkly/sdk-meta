---
id: java-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback
sdk: java-server-sdk
kind: reference
lang: java
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for Java.
validation:
  scaffold: java-server-sdk/scaffolds/java-syntax-only
---

```java
import com.launchdarkly.sdk.server.*;
import com.launchdarkly.sdk.server.integrations.*;

import java.net.URI;

URI relayUri = URI.create("http://my-relay-proxy:8030");
ServiceEndpointsBuilder relayEndpoints = Components.serviceEndpoints().relayProxy(relayUri);

LDConfig config = new LDConfig.Builder()
    .dataSystem(
        Components.dataSystem().custom()
            .initializers(
                DataSystemComponents.pollingInitializer()
                    .serviceEndpointsOverride(relayEndpoints),
                DataSystemComponents.pollingInitializer()
            )
            .synchronizers(
                DataSystemComponents.streamingSynchronizer()
                    .serviceEndpointsOverride(relayEndpoints),
                DataSystemComponents.streamingSynchronizer(),
                DataSystemComponents.pollingSynchronizer()
            )
    )
    .build();

LDClient client = new LDClient("YOUR_SDK_KEY", config);
```
