---
id: roku-client-sdk/sdk-docs/features/config/service-endpoint-configuration
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Service endpoint configuration example for Roku.
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
' for a legacy Roku application
config = LaunchDarklyConfig("example-mobile-key")

' for a SceneGraph Roku Application
config = LaunchDarklyConfig("example-mobile-key", CLIENT_SCENEGRAPH_NODE)

config.setStreamURI("https://your-relay-proxy.com:8030")
config.setAppURI("https://your-relay-proxy.com:8030")
config.setEventsURI("https://your-relay-proxy.com:8030")
```
