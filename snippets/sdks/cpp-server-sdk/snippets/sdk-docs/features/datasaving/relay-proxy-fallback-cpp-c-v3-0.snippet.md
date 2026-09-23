---
id: cpp-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback-cpp-c-v3-0
sdk: cpp-server-sdk
kind: reference
lang: c
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for C++ (server-side).
validation:
  scaffold: cpp-server-sdk/scaffolds/cpp-syntax-only
---

```c
const char* relay_url = "http://my-relay-proxy:8030";

LDServerFDv2PollingBuilder initializer_relay = LDServerFDv2PollingBuilder_New();
LDServerFDv2PollingBuilder_BaseURL(initializer_relay, relay_url);
LDServerFDv2PollingBuilder initializer_ld = LDServerFDv2PollingBuilder_New();

LDServerFDv2StreamingBuilder synchronizer_relay = LDServerFDv2StreamingBuilder_New();
LDServerFDv2StreamingBuilder_BaseURL(synchronizer_relay, relay_url);
LDServerFDv2StreamingBuilder synchronizer_ld = LDServerFDv2StreamingBuilder_New();
LDServerFDv2PollingBuilder synchronizer_poll = LDServerFDv2PollingBuilder_New();

LDServerFDv2Builder fdv2 = LDServerFDv2Builder_Custom();
LDServerFDv2Builder_Initializer_Polling(fdv2, initializer_relay);
LDServerFDv2Builder_Initializer_Polling(fdv2, initializer_ld);
LDServerFDv2Builder_Synchronizer_Streaming(fdv2, synchronizer_relay);
LDServerFDv2Builder_Synchronizer_Streaming(fdv2, synchronizer_ld);
LDServerFDv2Builder_Synchronizer_Polling(fdv2, synchronizer_poll);

LDServerConfigBuilder builder = LDServerConfigBuilder_New("YOUR_SDK_KEY");
LDServerConfigBuilder_DataSystem_FDv2(builder, fdv2);

LDServerConfig config;
LDStatus status = LDServerConfigBuilder_Build(builder, &config);

if (!LDStatus_Ok(status)) {
    /* an error occurred, config is not valid */
}
```
