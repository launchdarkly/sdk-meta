---
id: ios-client-sdk/sdk-docs/features/identify/identify-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Identify example for the iOS SDK v8.0+ (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
let newContext = try LDContextBuilder(key: "example-context-key").build().get();

// You can also call identify with a completion
LDClient.get()!.identify(context: newContext) {
    // Flags have been retrieved for the new context
}
```
