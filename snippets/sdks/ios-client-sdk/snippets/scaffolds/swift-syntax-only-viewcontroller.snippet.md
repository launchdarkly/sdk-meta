---
id: ios-client-sdk/scaffolds/swift-syntax-only-viewcontroller
sdk: ios-client-sdk
kind: scaffold
lang: swift
file: ViewController.swift
description: |
  Companion ViewController for the swift-syntax-only scaffold. The
  ios-client validator harness's `cp $SNIPPET_DIR/ViewController.swift`
  step expects this file to exist; the syntax-only path doesn't
  instantiate the controller (the EXAM-HELLO success line is
  emitted by AppDelegate's didFinishLaunchingWithOptions).
inputs: {}
---

```swift
import UIKit

class ViewController: UIViewController {
    @IBOutlet weak var featureFlagLabel: UILabel!
}
```
