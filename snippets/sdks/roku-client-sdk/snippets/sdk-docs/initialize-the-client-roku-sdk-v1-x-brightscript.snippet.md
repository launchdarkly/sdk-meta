---
id: roku-client-sdk/sdk-docs/initialize-the-client-roku-sdk-v1-x-brightscript
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: "Roku SDK v1.x (BrightScript) in section \"Initialize the client\""
# Bucket C: no Roku BrightScript validator. See _sdk-docs-port-notes.md.
---

```brightscript
' get a reference to the task
launchDarklyNode = m.top.findNode("my-node-name")

' create configuration
config = LaunchDarklyConfig("example-mobile-key", launchDarklyNode)

' create a user. You'll need this user later, but you can ignore it for now.
user = LaunchDarklyUser("example-user-key")

' create message port
messagePort = createObject("roMessagePort")

' initialize the client
LaunchDarklySGInit(config, user)

```
