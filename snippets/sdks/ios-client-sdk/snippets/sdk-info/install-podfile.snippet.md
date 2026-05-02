---
id: ios-client-sdk/sdk-info/install-podfile
sdk: ios-client-sdk
kind: install
lang: ruby
file: ios-client-sdk/install-podfile.txt
description: Install command for ios-client-sdk (podfile).
validation:
  runtime: ios-install
  env:
    INSTALL_KIND: podfile
---

```ruby
use_frameworks!
target 'YourTargetName' do
  pod 'LaunchDarkly', '~> 9.15'
end
```
