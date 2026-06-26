---
id: ios-client-sdk/sdk-docs/features/multienv/get-environment-objc
sdk: ios-client-sdk
kind: reference
lang: objectivec
description: Accessing a secondary environment client instance on iOS (Objective-C).

---

```objectivec
LDClient *coreInstance = [LDClient getWithEnvironment:@"core"];
// Variation determines whether or not a flag is enabled for a specific context
BOOL coreFlagValue = [coreInstance boolVariationForKey:@"core-example-flag-key" defaultValue:NO];
// allFlags produces a map of feature flag keys to their values
NSDictionary<NSString*, LDValue*> * allFlags = [coreInstance allFlags];
// track records actions end users take in your app
[coreInstance trackWithKey:@"MY_TRACK_EVENT_NAME" data:data];
```
