---
id: ios-client-sdk/sdk-docs/use-cocoapods-podfile
sdk: ios-client-sdk
kind: reference
lang: ruby
description: "Podfile in section \"Use CocoaPods\""
validation:
  runtime: ios-install
  # Filename the ios-install harness's `podfile` branch reads from
  # via `$BODY_FILE` (it copies the staged file to Podfile in a
  # workdir, then runs `pod install`). The Go validator stages every
  # runtime-based snippet to `frontmatter.file`, so the field is
  # required even though the harness picks up the body content too.
  file: Podfile
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
