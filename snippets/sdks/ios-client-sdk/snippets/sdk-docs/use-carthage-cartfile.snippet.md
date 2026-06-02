---
id: ios-client-sdk/sdk-docs/use-carthage-cartfile
sdk: ios-client-sdk
kind: reference
lang: swift
description: "Cartfile in section \"Use Carthage\""
# TODO(snippet-bug): body is Cartfile syntax
# (`github "owner/repo" ~> version`), not Swift — `~>` isn't a
# Swift operator and `github` isn't a Swift identifier. The source
# MDX tags this as `swift` but it's Carthage dependency-spec DSL.
# Fix in the snippet-bugs PR: re-tag (e.g. `text` or a custom
# `cartfile` lang) and skip syntax validation, or add a Cartfile
# parser path.
---

```swift
github "launchdarkly/ios-client" ~> 9.0
// optional observability plugin, requires iOS SDK v9.14+
github "launchdarkly/swift-launchdarkly-observability" ~> 1.0
```
