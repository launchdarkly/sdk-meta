---
id: android-client-sdk/scaffolds/init-runner-activity
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/MainActivity.kt
description: |
  Companion MainActivity for the Android init-runner scaffold.
  After MainApplication.onCreate has called LDClient.init(...),
  this activity reads the test flag via `LDClient.get()`, observes
  for changes, and writes the canonical EXAM-HELLO `feature flag
  evaluates to true` line into the TextView. The validator's
  Robolectric HelloAppTest drives this activity through its
  lifecycle and polls the TextView until it sees the line.

  Flag key is injected via `LAUNCHDARKLY_FLAG_KEY` (default
  `sample-feature` to match the EXAM-HELLO convention).
inputs: {}
---

```kotlin
package com.launchdarkly.hello_android

import android.os.Bundle
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import com.launchdarkly.sdk.android.LDClient

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val textView: TextView = findViewById(R.id.textview)

        val flagKey = System.getenv("LAUNCHDARKLY_FLAG_KEY") ?: "sample-feature"

        val client = LDClient.get()
        if (client == null) {
            textView.text = "scaffold: LDClient.get() returned null — init never ran"
            return
        }

        fun render(value: Boolean) {
            textView.text = if (value) {
                "feature flag evaluates to true"
            } else {
                "scaffold: flag evaluated to false"
            }
        }

        // Seed with the cached value, then register a listener so any
        // streaming-delivered update wins the assertion.
        render(client.boolVariation(flagKey, false))
        client.registerFeatureFlagListener(flagKey) {
            render(client.boolVariation(flagKey, false))
        }
    }
}
```
