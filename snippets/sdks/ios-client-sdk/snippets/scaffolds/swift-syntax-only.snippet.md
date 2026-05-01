---
id: ios-client-sdk/scaffolds/swift-syntax-only
sdk: ios-client-sdk
kind: scaffold
lang: swift
file: Snippet.swift
description: |
  Parse-only validator for iOS client SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: ios-client
  entrypoint: Snippet.swift
---

```swift
import Foundation

func _wrappee() {
{{ body }}
}

print("feature flag evaluates to true")
```
