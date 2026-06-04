---
id: roku-client-sdk/sdk-docs/features/config/app-config
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Application metadata configuration example for Roku.
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
' for a legacy Roku application
config = LaunchDarklyConfig("example-mobile-key")

' for a SceneGraph Roku Application
config = LaunchDarklyConfig("example-mobile-key", CLIENT_SCENEGRAPH_NODE)

' configure the application identifier and application version
config.setApplicationInfoValue("id", "authentication-service")
config.setApplicationInfoValue("version", "1.0.0")
```
