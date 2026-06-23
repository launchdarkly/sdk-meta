---
id: roku-client-sdk/sdk-docs/features/privateattrs/config
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Private attribute configuration for Roku (BrightScript).
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
' All attributes marked private
config = LaunchDarklyConfig("example-mobile-key", launchDarklyTaskNode)
config.setAllAttributesPrivate(true)

LaunchDarklySGInit(config, context)
client = LaunchDarklySG(launchDarklyTaskNode)

' Two attributes marked private
config = LaunchDarklyConfig("example-mobile-key", launchDarklyTaskNode)
config.addPrivateAttribute("email")
config.addPrivateAttribute("address")

LaunchDarklySGInit(config, context)
client = LaunchDarklySG(launchDarklyTaskNode)
```
