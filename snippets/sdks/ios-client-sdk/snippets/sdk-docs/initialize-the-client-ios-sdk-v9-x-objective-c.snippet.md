---
id: ios-client-sdk/sdk-docs/initialize-the-client-ios-sdk-v9-x-objective-c
sdk: ios-client-sdk
kind: reference
lang: objective-c
description: "iOS SDK v9.x (Objective-C) in section \"Initialize the client\""
# Bucket C: . See _sdk-docs-port-notes.md.
---

```objectivec
  LDConfig *config = [[LDConfig alloc] initWithMobileKey:@"example-mobile-key" autoEnvAttributes:AutoEnvAttributesEnabled];

  // You'll need this context later, but you can ignore it for now.
  LDContextBuilder *builder = [[LDContextBuilder alloc] initWithKey:@"example-context-key"];
  [builder kindWithKind:@"user"];
  [builder nameWithName:@"Sandy"];
  ContextBuilderResult *result = [builder build];

  [LDClient startWithConfiguration:config context:context startWaitSeconds:5.0 completion:^(BOOL timedOut) {
    if(timedOut) {
        // Client may not have the most recent flags for the configured context
    } else {
        // Client has received flags for the configured context
    }
  }];
```
