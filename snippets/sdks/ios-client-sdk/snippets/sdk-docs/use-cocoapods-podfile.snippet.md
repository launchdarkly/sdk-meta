---
id: ios-client-sdk/sdk-docs/use-cocoapods-podfile
sdk: ios-client-sdk
kind: reference
lang: ruby
description: "Podfile in section \"Use CocoaPods\""
# Top-level `file:` — the Go validator stages every runtime-based
# snippet to `frontmatter.file`. The ios-install harness's `podfile`
# branch then copies it to `Podfile` and runs `pod install`.
file: ios-client-sdk/sdk-docs/use-cocoapods-podfile.Podfile
validation:
  runtime: ios-install
  env:
    INSTALL_KIND: podfile
---

```ruby
use_frameworks!
target 'YourTargetName' do
  pod 'LaunchDarkly', '~> 9.0'
  # optional observability plugin, requires iOS SDK v9.14+
  pod 'LaunchDarklyObservability', '~> 1.0'
end
```
