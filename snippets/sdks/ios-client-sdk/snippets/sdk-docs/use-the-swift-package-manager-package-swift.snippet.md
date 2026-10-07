---
id: ios-client-sdk/sdk-docs/use-the-swift-package-manager-package-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "Package.swift in section \"Use the Swift Package Manager\""
file: ios-client-sdk/use-the-swift-package-manager-package-swift.txt
validation:
  runtime: ios-install
  env:
    INSTALL_KIND: swift-package
---

```swift
//...
    dependencies: [
        .package(url: "https://github.com/launchdarkly/ios-client-sdk.git", .upToNextMinor(from: "11.6.2")),
        // optional observability plugin, requires iOS SDK v11.5+
        .package(url: "https://github.com/launchdarkly/swift-launchdarkly-observability.git", .upToNextMajor(from: "0.56.0")),
    ],
    targets: [
        .target(
            name: "YOUR_TARGET",
            dependencies: ["LaunchDarkly"]
        )
    ],
//...
```
