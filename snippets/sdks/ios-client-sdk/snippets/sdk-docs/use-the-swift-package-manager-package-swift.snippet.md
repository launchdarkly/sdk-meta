---
id: ios-client-sdk/sdk-docs/use-the-swift-package-manager-package-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "Package.swift in section \"Use the Swift Package Manager\""
---

```swift
//...
    dependencies: [
        .package(url: "https://github.com/launchdarkly/ios-client-sdk.git", .upToNextMinor("9.0.0")),
        // optional observability plugin, requires iOS SDK v9.14+
        .package(url: "https://github.com/launchdarkly/swift-launchdarkly-observability.git", .upToNextMajor("1.0.0")),
    ],
    targets: [
        .target(
            name: "YOUR_TARGET",
            dependencies: ["LaunchDarkly"]
        )
    ],
//...
```
