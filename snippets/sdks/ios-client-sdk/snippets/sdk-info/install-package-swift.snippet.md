---
id: ios-client-sdk/sdk-info/install-package-swift
sdk: ios-client-sdk
kind: install
lang: swift
file: ios-client-sdk/install-package-swift.txt
description: Install command for ios-client-sdk (package-swift).
---

```swift
//...
    dependencies: [
        .package(url: "https://github.com/launchdarkly/ios-client-sdk.git", .upToNextMajor(from: "9.15.0")),
   ],
    targets: [
        .target(
            name: "YOUR_TARGET",
            dependencies: ["LaunchDarkly"]
        )
    ],
//...
```
