---
id: ios-client-sdk/sdk-docs/features/identify/identify-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: Identify example for the iOS SDK v9.4+ (Swift).
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
let newContext = try LDContextBuilder(key: "example-context-key").build().get()

LDClient.get()!.identify(context: newContext, timeout: 5) { result in
    switch result {
    case .complete:
        // Flags have been retrieved for the new context
        break
    case .timeout:
        // The identify request is still in progress
        break
    case .shed:
        // A later identify call replaced this one
        break
    case .error:
        // The identify request failed
        break
    }
}
```
