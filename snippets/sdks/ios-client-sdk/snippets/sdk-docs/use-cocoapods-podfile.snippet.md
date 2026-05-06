---
id: ios-client-sdk/sdk-docs/use-cocoapods-podfile
sdk: ios-client-sdk
kind: reference
lang: ruby
description: "Podfile in section \"Use CocoaPods\""
# Bucket C: . See _sdk-docs-port-notes.md.
---

```ruby
use_frameworks!
target 'YourTargetName' do
  pod 'LaunchDarkly', '~> 9.0'
  # optional observability plugin, requires iOS SDK v9.14+
  pod 'LaunchDarklyObservability', '~> 1.0'
end
```
