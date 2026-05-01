---
id: roku-client-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-private-attributes-2-0-syntax-two-attributes-marked-private
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: "2.0 syntax, two attributes marked private in section \"Understanding changes to private attributes\""
---

```brightscript
config = LaunchDarklyConfig("example-mobile-key", launchDarklyTaskNode)
config.addPrivateAttribute("email")
config.addPrivateAttribute("address")

LaunchDarklySGInit(config, context)
client = LaunchDarklySG(launchDarklyTaskNode)
```
