---
id: android-client-sdk/getting-started/main-activity
sdk: android-client-sdk
kind: hello-world
lang: kotlin
file: app/src/main/java/com/launchdarkly/hello_android/MainActivity.kt
description: MainActivity that observes the flag and renders the value.
inputs:
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: main-activity
# Validator pending — Android validation needs setup-android + a Linux
# emulator boot, slow but feasible. Deferred.
---

Open the file `MainActivity.kt` and add the following code:

```kotlin
package com.launchdarkly.hello_android

import android.os.Bundle
import android.view.View
import android.widget.TextView
import androidx.appcompat.app.AlertDialog
import androidx.appcompat.app.AppCompatActivity
import com.launchdarkly.hello_android.MainApplication.Companion.LAUNCHDARKLY_MOBILE_KEY
import com.launchdarkly.sdk.android.LDClient

class MainActivity : AppCompatActivity() {

    // Set BOOLEAN_FLAG_KEY to the feature flag key you want to evaluate.
    val BOOLEAN_FLAG_KEY = "{{ featureKey }}"

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val textView : TextView = findViewById(R.id.textview)
        val fullView : View = window.decorView

        if (LAUNCHDARKLY_MOBILE_KEY == "example-mobile-key") {
            val builder = AlertDialog.Builder(this)
            builder.setMessage("LAUNCHDARKLY_MOBILE_KEY was not customized for this application.")
            builder.create().show()
        }

        val client = LDClient.get()
        val flagValue = client.boolVariation(BOOLEAN_FLAG_KEY, false)

        // to get the variation the SDK has cached
        textView.text = getString(
            R.string.flag_evaluated,
            BOOLEAN_FLAG_KEY,
            flagValue.toString()
        )

        // Style the display
        textView.setTextColor(resources.getColor(R.color.colorText))
        if(flagValue) {
            fullView.setBackgroundColor(resources.getColor(R.color.colorBackgroundTrue))
        } else {
            fullView.setBackgroundColor(resources.getColor(R.color.colorBackgroundFalse))
        }

        // to register a listener to get updates in real time
        client.registerFeatureFlagListener(BOOLEAN_FLAG_KEY) {
            val changedFlagValue = client.boolVariation(BOOLEAN_FLAG_KEY, false)
            textView.text = getString(
                R.string.flag_evaluated,
                BOOLEAN_FLAG_KEY,
                changedFlagValue.toString()
            )
            if(changedFlagValue) {
                fullView.setBackgroundColor(resources.getColor(R.color.colorBackgroundTrue))
            } else {
                fullView.setBackgroundColor(resources.getColor(R.color.colorBackgroundFalse))
            }
        }
    }
}
```
