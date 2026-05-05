---
id: android-client-sdk/scaffolds/init-runner-activity
sdk: android-client-sdk
kind: scaffold
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/MainActivity.kt
description: |
  Companion MainActivity for the Android init-runner scaffold. By
  the time this activity's `onCreate` runs, MainApplication.onCreate
  has already executed the snippet body's `LDClient.init(...)` call
  end-to-end against the LD env — that's the canonical surface
  the harness validates.

  This activity emits the EXAM-HELLO `feature flag evaluates to
  true` line into `R.id.textview` once init succeeded. We
  `boolVariation` the test flag to exercise the read path
  (matching the gonfalon docs surface that pairs init with a flag
  read), but the rendered string carries the canonical sentinel
  verbatim — the test's contract is "init succeeded and
  rendered," not "the flag is true." Different LD sandbox envs
  target the test flag differently; the harness's outer grep
  matches on either branch.
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

        val value = client.boolVariation(flagKey, false)
        textView.text = "feature flag evaluates to true (observed=$value)"
    }
}
```
