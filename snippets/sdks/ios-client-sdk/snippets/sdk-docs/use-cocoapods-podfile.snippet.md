---
id: ios-client-sdk/sdk-docs/use-cocoapods-podfile
sdk: ios-client-sdk
kind: reference
lang: ruby
description: "Podfile in section \"Use CocoaPods\""
file: ios-client-sdk/use-cocoapods-podfile.txt
validation:
  runtime: ios-install
  env:
    INSTALL_KIND: podfile
---

```ruby
use_frameworks!
target 'YourTargetName' do
  pod 'LaunchDarkly', '~> 11.0'
  # optional observability plugin, requires iOS SDK v11.5+
  pod 'LaunchDarklyObservability', '~> 0.56'
end
```
