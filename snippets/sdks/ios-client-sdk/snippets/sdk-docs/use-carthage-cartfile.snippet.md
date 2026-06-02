---
id: ios-client-sdk/sdk-docs/use-carthage-cartfile
sdk: ios-client-sdk
kind: reference
lang: swift
description: "Cartfile in section \"Use Carthage\""
# TODO(validate): . See _sdk-docs-port-notes.md.
validation:
  scaffold: ios-client-sdk/scaffolds/swift-syntax-only
---

```swift
github "launchdarkly/ios-client" ~> 9.0
// optional observability plugin, requires iOS SDK v9.14+
github "launchdarkly/swift-launchdarkly-observability" ~> 1.0
```
