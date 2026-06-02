---
id: ios-client-sdk/sdk-docs/use-the-swift-package-manager-package-swift
sdk: ios-client-sdk
kind: reference
lang: swift
description: "Package.swift in section \"Use the Swift Package Manager\""
# TODO(snippet-bug): body's
# `swift-launchdarkly-observability.git`, .upToNextMajor("1.0.0")`
# references a 1.x major-version range, but the observability
# package only has 0.x tags published. Swift Package Manager's
# resolver would fail to find a compatible version. Top-level
# `file:` is also missing — the Go validator stages every
# runtime-based snippet to `frontmatter.file`. Fix in the
# snippet-bugs PR: drop the observability line (or update once a
# 1.x tag exists) and add `file: Package.swift`.
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
