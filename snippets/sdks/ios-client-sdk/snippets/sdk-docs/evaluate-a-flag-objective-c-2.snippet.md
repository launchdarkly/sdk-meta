---
id: ios-client-sdk/sdk-docs/evaluate-a-flag-objective-c-2
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "Objective-C in section \"Evaluate a flag\""
# TODO(validate): . See _sdk-docs-port-notes.md.
---

```objectivec
BOOL showFeature = [client boolVariationForKey:@"example-flag-key" defaultValue:NO];
if (showFeature) {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
