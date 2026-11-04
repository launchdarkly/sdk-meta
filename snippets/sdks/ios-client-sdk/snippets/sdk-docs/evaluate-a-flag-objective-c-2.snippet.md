---
id: ios-client-sdk/sdk-docs/evaluate-a-flag-objective-c-2
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "Objective-C in section \"Evaluate a flag\""
# TODO(validate): swift-syntax-only compiles Swift only. To validate, add an Objective-C scaffold
# (or a harness branch on lang:) and a project.yml with an Objective-C target.
---

```objectivec
BOOL showFeature = [client boolVariationForKey:@"example-flag-key" defaultValue:NO];
if (showFeature) {
    // Application code to show the feature
} else {
    // The code to run if the feature is off
}
```
