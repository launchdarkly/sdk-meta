---
id: roku-client-sdk/sdk-docs/features/evaluating/evaluating
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: Flag evaluation example for Roku.
validation:
  scaffold: roku-client-sdk/scaffolds/roku-syntax-only

---

```brightscript
' typed variations
myInt = launchDarkly.intVariation("example-flag-key", 123)

myBool = launchDarkly.boolVariation("example-flag-key", false)

myString = launchDarkly.stringVariation("example-flag-key", "hello world!")

myObjectOrArray = launchDarkly.jsonVariation("example-flag-key", {"value": 123})

' generic variation
myValue = launchDarkly.variation("example-flag-key", false)
```
