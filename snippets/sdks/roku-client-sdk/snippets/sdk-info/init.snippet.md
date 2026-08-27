---
id: roku-client-sdk/sdk-info/init
sdk: roku-client-sdk
kind: init
lang: brightscript
file: roku-client-sdk/init.txt
description: |
  Client initialization snippet for roku-client-sdk (SceneGraph, v2 API).
  Parse-checked via the roku-syntax-only scaffold (brighterscript parser;
  no Roku device in CI).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only
---

```brightscript
' get a reference to the LaunchDarklyTask node in your scene
launchDarklyNode = m.top.findNode("launchDarkly")

' create a configuration with your mobile key
config = LaunchDarklyConfig("YOUR_MOBILE_KEY", launchDarklyNode)

' create a context
context = LaunchDarklyCreateContext({"key": "example-user-key", "kind": "user"})

' initialize the client
LaunchDarklySGInit(config, context)
```
