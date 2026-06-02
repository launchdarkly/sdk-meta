---
id: ios-client-sdk/sdk-docs/use-cocoapods-podfile
sdk: ios-client-sdk
kind: reference
lang: ruby
description: "Podfile in section \"Use CocoaPods\""
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
