---
id: flutter-client-sdk/getting-started/podfile-platform
sdk: flutter-client-sdk
kind: manifest-fragment
lang: ruby
description: ios/Podfile platform line.
ld-application:
  slot: podfile-platform
---

Ensure that `ios/Podfile` specifies a minimum deployment target of at least 10.0.

```ruby
platform :ios, '10.0'
```
