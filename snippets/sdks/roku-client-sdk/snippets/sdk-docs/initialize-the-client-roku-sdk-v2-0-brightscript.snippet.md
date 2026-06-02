---
id: roku-client-sdk/sdk-docs/initialize-the-client-roku-sdk-v2-0-brightscript
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: "Roku SDK v2.0 (BrightScript) in section \"Initialize the client\""
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only
---

```brightscript
' get a reference to the task
launchDarklyNode = m.top.findNode("my-node-name")

' create configuration
config = LaunchDarklyConfig("example-mobile-key", launchDarklyNode)

' create a context. You'll need this context later, but you can ignore it for now.
context = LaunchDarklyCreateContext({"key": "example-user-key", "kind": "user"})

' create message port
messagePort = createObject("roMessagePort")

' initialize the client
LaunchDarklySGInit(config, context)

```
